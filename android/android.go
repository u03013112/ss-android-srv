package android

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	pb "github.com/u03013112/ss-pb/android"
)

// Srv ：服务
type Srv struct{}

// Login :
func (s *Srv) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	user := getOrCreateUserByUUID(in.Uuid)
	token, _ := uuid.NewV4()
	user.updateToken(token.String())
	return &pb.LoginReply{
		Token:       user.Token,
		ExpiresDate: user.ExpireDate.Unix(),
		Total:       user.TotalRxTraffic,
		Used:        user.UsedRxTraffic,
	}, nil
}

// GetConfig :
func (s *Srv) GetConfig(ctx context.Context, in *pb.GetConfigRequest) (*pb.GetConfigReply, error) {
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
	if user.TotalRxTraffic < user.UsedRxTraffic {
		return &pb.GetConfigReply{
			Error: "流量不足，请及时续费",
		}, nil
	}
	ret := new(pb.GetConfigReply)
	if config, err := grpcGetConfig(); err != nil {
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
	ret.Total = user.TotalRxTraffic
	use := int64(0)
	if in.Rx >= user.BaseRxTraffic {
		use = in.Rx - user.BaseRxTraffic
	} else {
		use = in.Rx
	}
	user.BaseRxTraffic = in.Rx
	user.UsedRxTraffic += use
	ret.Used = user.UsedRxTraffic
	if user.UsedRxTraffic >= user.TotalRxTraffic {
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
	}

	user.update()

	bill := new(Bill)
	bill.UserID = int64(user.ID)
	bill.ProductionID = int64(p.ID)
	bill.record()

	ret := pb.BuyTestReply{
		ExpiresDate: user.ExpireDate.Unix(),
		Total:       user.TotalRxTraffic,
		Used:        user.UsedRxTraffic,
	}

	return &ret, nil
}

// GetGoogleAd :
func (s *Srv) GetGoogleAd(ctx context.Context, in *pb.GetGoogleAdRequest) (*pb.GetGoogleAdReply, error) {
	ret := new(pb.GetGoogleAdReply)
	ret.Id = getGoogleAd()
	return ret,nil
}