package apiserver

import "google.golang.org/grpc"

type grpcAPIServer struct {
	*grpc.Server
	address string
}
