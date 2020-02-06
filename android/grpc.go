package android

import (
	"context"
	"log"
	"time"

	config "github.com/u03013112/ss-pb/config"
	"google.golang.org/grpc"
)

const (
	configAddress = "config:50001"
)

func grpcGetConfig(lineID int64) (*config.GetSSConfigReply, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(configAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := config.NewSSConfigClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetSSConfig(ctx, &config.GetSSConfigRequest{
		Token:  "u03013112",
		LineID: lineID,
	})
	if err != nil {
		log.Printf("could not GetSSConfig: %v", err)
		return r, err
	}
	return r, nil
}

func GrpcSetPasswd(passwd string) error {
	conn, err := grpc.Dial(configAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := config.NewSSConfigClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SetPasswd(ctx, &config.SetPasswdRequest{Passwd: passwd})
	if err != nil {
		log.Printf("could not SetPasswd: %v", err)
		return err
	}
	return nil
}
