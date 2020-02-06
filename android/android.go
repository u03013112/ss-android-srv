package android

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	pb "github.com/u03013112/ss-pb/android"
	"google.golang.org/grpc/metadata"
)

// Srv ：服务
type Srv struct{}

// DebugJSON : DebugJSON
func DebugJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[%s]", data)
}

// Login :
func (s *Srv) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	ip := "unknow"
	version := "too old"
	mi, _ := metadata.FromIncomingContext(ctx)
	// DebugJSON(mi)
	if x, ok := mi["x-forwarded-for"]; ok == true {
		ipList := x[0]
		ip = strings.Split(ipList, ",")[0]
	}

	user := getOrCreateUserByUUID(in.Uuid)
	token, _ := uuid.NewV4()

	if in.Version != "" {
		version = in.Version
	}

	user.updateUserInfo(token.String(), ip, version)

	log := new(LoginLog)
	log.UserID = int64(user.ID)
	log.IP = ip
	log.record()

	return &pb.LoginReply{
		Token:       user.Token,
		ExpiresDate: user.ExpireDate.Unix(),
		Total:       user.TotalRxTraffic,
		Used:        user.UsedRxTraffic + user.UsedTxTraffic,
	}, nil
}

// GetConfig :
func (s *Srv) GetConfig(ctx context.Context, in *pb.GetConfigRequest) (*pb.GetConfigReply, error) {
	return &pb.GetConfigReply{
		Error: "版本过旧，请及时升级！",
	}, nil

	// 接口作废
	user, err := getUserByToken(in.Token)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	if user.ExpireDate.Unix() < time.Now().Unix() {
		return &pb.GetConfigReply{
			Error: "已过期，请及时续费",
		}, nil
	}
	if user.TotalRxTraffic < user.UsedRxTraffic+user.UsedTxTraffic {
		return &pb.GetConfigReply{
			Error: "流量不足，请及时续费",
		}, nil
	}
	ret := new(pb.GetConfigReply)
	if config, err := grpcGetConfig(in.LineID); err != nil {
		fmt.Print(err.Error())
		return ret, err
	} else {
		ret.IP = config.IP
		ret.Port = config.Port
		ret.Method = config.Method
		ret.Passwd = config.Passwd
	}
	return ret, nil
}

// GetConfigNew :
func (s *Srv) GetConfigNew(ctx context.Context, in *pb.GetConfigRequest) (*pb.GetConfigNewReply, error) {
	ret := new(pb.GetConfigNewReply)
	ret0 := new(pb.GetConfigReply)

	user, err := getUserByToken(in.Token)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	if user.ExpireDate.Unix() < time.Now().Unix() {
		ret0.Error = "已过期，请及时续费"
	}
	if user.TotalRxTraffic < user.UsedRxTraffic+user.UsedTxTraffic {
		ret0.Error = "流量不足，请及时续费"
	}
	config, err := grpcGetConfig(in.LineID)
	if err != nil {
		fmt.Print(err.Error())
		return ret, err
	}
	if ret0.Error == "" {
		ret0.IP = config.IP
		ret0.Port = config.Port
		ret0.Method = config.Method
		ret0.Passwd = config.Passwd
	}
	j, _ := json.Marshal(ret0)
	ret.Config = encode(string(j))
	return ret, nil
}

// GetConfigV1 :
func (s *Srv) GetConfigV1(ctx context.Context, in *pb.GetConfigRequest) (*pb.GetConfigNewReply, error) {
	ret := new(pb.GetConfigNewReply)
	ret0 := new(pb.GetConfigReply)

	user, err := getUserByToken(in.Token)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	if user.ExpireDate.Unix() < time.Now().Unix() {
		ret0.Error = "已过期，请及时续费"
	}
	if user.TotalRxTraffic < user.UsedRxTraffic+user.UsedTxTraffic {
		ret0.Error = "流量不足，请及时续费"
	}
	config, err := grpcGetConfig(in.LineID)
	if err != nil {
		fmt.Print(err.Error())
		return ret, err
	}
	if ret0.Error == "" {
		ret0.IP = config.IP
		ret0.Port = config.Port
		ret0.Method = config.Method
		ret0.Passwd = encode4Passwd(config.Passwd)
	}
	j, _ := json.Marshal(ret0)
	ret.Config = encode(string(j))
	return ret, nil
}

func encode4Passwd(passwd string) string {
	p := []byte(passwd)
	for i := 0; i < len(p); i++ {
		p[i] = p[i] + 3
	}
	passwd = string(p)
	return passwd
}

func encode(str string) string {
	var i = 'a'
	var ret = base64.StdEncoding.EncodeToString([]byte(str))
	for i = 'a'; i <= 'z'; i++ {
		ret = strings.Replace(ret, fmt.Sprintf("%c", i), fmt.Sprintf(".%c", i-1), -1)
	}
	return ret
}

// Keepalive :
func (s *Srv) Keepalive(ctx context.Context, in *pb.KeepaliveRequest) (*pb.KeepaliveReply, error) {
	ret := new(pb.KeepaliveReply)
	user, err := getUserByToken(in.Token)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	ret.ExpiresDate = user.ExpireDate.Unix()
	ret.NeedStop = false
	if user.ExpireDate.Unix() < time.Now().Unix() {
		ret.NeedStop = true
	}
	// 总量，虽然名字叫Rx，但是目前是双方向的，这个统计方式需要再尝试。
	ret.Total = user.TotalRxTraffic

	rxUse := int64(0)
	if in.Rx >= user.BaseRxTraffic {
		rxUse = in.Rx - user.BaseRxTraffic
	} else {
		rxUse = in.Rx
	}
	user.BaseRxTraffic = in.Rx
	user.UsedRxTraffic += rxUse

	txUse := int64(0)
	if in.Tx >= user.BaseTxTraffic {
		txUse = in.Tx - user.BaseTxTraffic
	} else {
		txUse = in.Tx
	}
	user.BaseTxTraffic = in.Tx
	user.UsedTxTraffic += txUse

	ret.Used = user.UsedRxTraffic + user.UsedTxTraffic
	if user.UsedRxTraffic+user.UsedTxTraffic >= user.TotalRxTraffic {
		ret.NeedStop = true
	}

	user.update()
	return ret, nil
}

// GetProdectionList :
func (s *Srv) GetProdectionList(ctx context.Context, in *pb.GetProdectionListRequest) (*pb.GetProdectionListReply, error) {
	ret := new(pb.GetProdectionListReply)
	pList, err := (new(Production)).findAll()
	for _, p := range pList {
		if p.Hiden != 0 {
			continue
		}
		pro := pb.Prodection{
			ID:          int64(p.ID),
			Time:        p.Time,
			Total:       p.Total,
			Price:       p.Price,
			Description: p.Description,
		}
		ret.ProdectionList = append(ret.ProdectionList, &pro)
	}
	return ret, err
}

// BuyTest :
func (s *Srv) BuyTest(ctx context.Context, in *pb.BuyTestRequest) (*pb.BuyTestReply, error) {
	user, err := getUserByToken(in.Token)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	p := new(Production)
	p.ID = uint(in.ProdectionID)
	if err := p.getByID(); err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	if user.ExpireDate.Unix() > time.Now().Unix() { //有效期内
		user.ExpireDate = user.ExpireDate.AddDate(0, 0, int(p.Time))
		user.TotalRxTraffic = user.TotalRxTraffic + p.Total
	} else {
		user.ExpireDate = time.Now().AddDate(0, 0, int(p.Time))
		user.TotalRxTraffic = p.Total
		user.UsedRxTraffic = 0
		user.UsedTxTraffic = 0
	}

	user.update()

	bill := new(Bill)
	bill.UserID = int64(user.ID)
	bill.ProductionID = int64(p.ID)
	bill.record()

	ret := pb.BuyTestReply{
		ExpiresDate: user.ExpireDate.Unix(),
		Total:       user.TotalRxTraffic,
		Used:        user.UsedRxTraffic + user.UsedTxTraffic,
	}

	return &ret, nil
}

// GetGoogleAd :
func (s *Srv) GetGoogleAd(ctx context.Context, in *pb.GetGoogleAdRequest) (*pb.GetGoogleAdReply, error) {
	ret := new(pb.GetGoogleAdReply)
	ret.Id = getGoogleAd()
	return ret, nil
}

// GetGoogleAdFree :
func (s *Srv) GetGoogleAdFree(ctx context.Context, in *pb.GetGoogleAdRequest) (*pb.GetGoogleAdReply, error) {
	ret := new(pb.GetGoogleAdReply)
	ret.Id = getGoogleAdFree()
	return ret, nil
}

// GetUserInfo :
func (s *Srv) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoReply, error) {
	ret := &pb.GetUserInfoReply{
		Status: "normal",
	}
	user, err := getUserByToken(in.Token)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	if user.ExpireDate.Unix() < time.Now().Unix() {
		ret.Status = "expired"
	}
	if user.TotalRxTraffic < user.UsedRxTraffic+user.UsedTxTraffic {
		ret.Status = "outOfTraffic"
	}
	return ret, nil
}
