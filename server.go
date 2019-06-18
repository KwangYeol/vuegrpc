package main

import (
    "fmt"
    "log"
    "net"

    // "github.com/kwangyeol/vuegrpc/todo"
    "./todo"
    "google.golang.org/grpc"
)

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9999))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := todo.Server {}
    grpcServer := grpc.NewServer()
    // attach the todo service to the server
    todo.RegisterTodoServiceServer(grpcServer, &s)
    
    if err := grpcServer.Serve(lis); err != nil {
        log.Printf("444")
        log.Fatalf("failed to serve: %s", err)
    } else {
        log.Printf("Server started successfully")
    }
}