package main

import (
	"log"
	"net"

	"github.com/u03013112/ss-android/android"
	"github.com/u03013112/ss-android/schedule"
	pb "github.com/u03013112/ss-pb/android"
	"google.golang.org/grpc"
)

const (
	port = ":50003"
)

func main() {
	android.InitDB()
	schedule.Init()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen %s", port)
	s := grpc.NewServer()
	pb.RegisterAndroidServer(s, &android.Srv{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
