package main

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/abdybaevae/url-shortener/gateway"
	"github.com/abdybaevae/url-shortener/insecure"
	keys_service "github.com/abdybaevae/url-shortener/pkg/services/keys"
	links_service "github.com/abdybaevae/url-shortener/pkg/services/links"

	pbExample "github.com/abdybaevae/url-shortener/proto"
	"github.com/abdybaevae/url-shortener/server"
)

func main() {
	// init services
	keyService := keys_service.NewService()
	linkService := links_service.NewLinkService(keyService)

	//
	backend := server.NewBackend(linkService)

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	pbExample.RegisterLinkServiceServer(s, backend)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
