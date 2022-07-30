package main

import (
	"context"
	"log"
	"net"

	"github.com/rohanray/modular-microsvc-dapr-todo/db-go-ent-grpc/ent"
	"github.com/rohanray/modular-microsvc-dapr-todo/db-go-ent-grpc/ent/proto/entpb"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	client, err := ent.Open("mysql", "fcz9sztdmsw3:pscale_pw_srCWXLKbsYPR4vtRc95F55XEjNVv28XNrQ5sJ_qWuLk@tcp(p2mxyyzutxse.ap-southeast-2.psdb.cloud)/dapr_todo?tls=true&parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize the generated User service.
	svc := entpb.NewUserService(client)

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()

	// Register the User service with the server.
	entpb.RegisterUserServiceServer(server, svc)

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
