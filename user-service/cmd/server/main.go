package main

import (
	"context"
	"log"
	"net"

	pb "github.com/yesiamanerd/flight-reservation-golang/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// generated in user_grpc.pb.go:
// type UserServiceServer interface {
//     Signup(context.Context, *pb.SignupRequest) (*pb.AuthResponse, error)
//     Login (context.Context, *pb.LoginRequest) (*pb.AuthResponse, error)
// }

// So in your file:

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.AuthResponse, error) {
	// TODO: insert into database, hash password, etc.
	log.Printf("Signup received: username=%s", req.Username)
	return &pb.AuthResponse{
		UserId: 1,
		Roles:  []string{"user"},
		Status: "active",
	}, nil
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	log.Printf("Login received: username=%s", req.Username)
	return &pb.AuthResponse{
		UserId: 1,
		Roles:  []string{"user"},
		Status: "active",
	}, nil
}

func main() {
	// 1) Listen on TCP port 50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2) Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// 3) Register your UserService implementation with it
	pb.RegisterUserServiceServer(grpcServer, &server{})

	// ────────────────────────────────────────────────────
	// 4) Enable gRPC Server Reflection here (for grpcurl & other tools)
	reflection.Register(grpcServer)
	// ────────────────────────────────────────────────────

	// 5) Log and start serving
	log.Println("UserService listening on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
