package main

import (
	"context"
	pb "grpc_prtc/usermgmt"
	"log"
)

var (
	port = ":30000"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateUser(ctx context.Context, in pb.NewUser) {
	log.Printf("Received: %v", in.GetName())
	return &pb.NewUser{Message: } 
}