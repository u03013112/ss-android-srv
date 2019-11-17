package android

import (
	"context"
	"errors"
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
	if user.ExpireDate.Unix() > time.Now().Unix() {
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
	return &pb.GetConfigReply{}, errors.New("ExpireDate")
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
	if user.ExpireDate.Unix() > time.Now().Unix() {
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
	user.UsedRxTraffic = use + user.UsedRxTraffic
	ret.Used = user.UsedRxTraffic
	user.update()
	return ret, nil
}
