package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc_prtc/usermgmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = ":30000"
)

func main() {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	//Contact the server & get the response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Alice"] = 43
	new_users["Bob"] = 31

	for username, age := range new_users {
		r, err := c.CreateUser(ctx, &pb.NewUser{Username: username, Age: age})
		if err != nil {
			log.Fatalf("could not create a user: %v", err)
		}
		log.Printf(`Created user: ID: %d USERNAME: %s AGE: %d`, r.GetId(), r.GetUsername(), r.GetAge())

	}
	params := &pb.GetUsersParams{}
	r, err := c.GetUsers(ctx, params)
	if err != nil {
		log.Fatalf("could not retrieve users: %v", err)
	}

	log.Print("\nUSER LIST: \n")
	fmt.Printf("aaaaaaaaaaaaaaaaaand the r.GetUsers() have this data: %v\n", r.GetUsers())
}
