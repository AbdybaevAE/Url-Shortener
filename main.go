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
	"github.com/abdybaevae/url-shortener/pkg/database"
	"github.com/abdybaevae/url-shortener/pkg/migrations"

	algo_repo "github.com/abdybaevae/url-shortener/pkg/repos/algo"
	algo_service "github.com/abdybaevae/url-shortener/pkg/services/algo"

	key_repo "github.com/abdybaevae/url-shortener/pkg/repos/key"
	key_service "github.com/abdybaevae/url-shortener/pkg/services/key"

	link_repo "github.com/abdybaevae/url-shortener/pkg/repos/link"
	link_service "github.com/abdybaevae/url-shortener/pkg/services/link"

	pbExample "github.com/abdybaevae/url-shortener/proto"
	"github.com/abdybaevae/url-shortener/server"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	db, err := database.Conn()
	if err != nil {
		log.Fatalln("cannot connect to database ", db)
	}

	// init services
	algoRepo := algo_repo.New()
	algoService := algo_service.New(algoRepo)

	keyRepo := key_repo.New()
	keyService := key_service.New(keyRepo, algoService)

	linkRepo := link_repo.New()
	linkService := link_service.New(linkRepo, keyService)

	if err := migrations.Run(algoService); err != nil {
		log.Fatal("cannot migrate", err)
	}

	//
	backend := server.NewBackend(linkService)

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!

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
