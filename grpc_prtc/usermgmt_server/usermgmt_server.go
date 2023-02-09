package main

import (
	"context"

	"log"
	"math/rand"
	"net"

	pb "grpc_prtc/usermgmt"

	"google.golang.org/grpc"
)

var (
	port = ":30000"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		user_list: &pb.Userlist{},
	}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.Userlist
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *UserManagementServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetUsername())
	var user_id int32 = int32(rand.Intn(100))
	created_user := &pb.User{Id: user_id, Username: in.GetUsername(), Age: in.GetAge()}
	s.user_list.Users = append(s.user_list.Users, created_user)

	return created_user, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.Userlist, error) {

	return s.user_list, nil
}

func (s *UserManagementServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	log.Printf("Received %v User to update", in.GetName())

	// arg := &pb.UpdateUserRequest{Id: in.GetId(), Name: in.GetName(), Age: in.GetAge()}

	// user, err := s.UpdateUser(ctx, arg)
	// if err != nil {
	// 	panic(err)
	// }

	// s.user_list.Users = append(s.user_list.Users, updated_user)

	// fmt.Println(Map(s.user_list.Users, in.GetName()))
	// rsp := &pb.UpdateUserResponse{
	// 	User: convertUser(user)
	// }

	updated_user := &pb.UpdateUserResponse{}

	// {Id: in.GetId(), Name: in.GetName(), Age: in.GetAge()}
	return updated_user, nil
}

func main() {
	var user_mgmt_server *UserManagementServer = NewUserManagementServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
