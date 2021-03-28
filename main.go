package main

import (
	"flag"
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

	num_repo "github.com/abdybaevae/url-shortener/pkg/repos/number"
	num_srv "github.com/abdybaevae/url-shortener/pkg/services/number"

	link_repo "github.com/abdybaevae/url-shortener/pkg/repos/link"
	link_service "github.com/abdybaevae/url-shortener/pkg/services/link"

	pbExample "github.com/abdybaevae/url-shortener/proto"
	"github.com/abdybaevae/url-shortener/server"
)

func main() {

	algorithm := flag.String("algorithm", string(algo_service.BASE_62), "Default algorithm name that will be used to pre generate keys")
	flag.Parse()
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	db, err := database.Conn()
	if err != nil {
		log.Fatalln("cannot connect to database ", db)
	}
	if err := migrations.Run(db); err != nil {
		log.Fatal("cannot migrate", err)
	}

	// init services
	numRepo := num_repo.New(db)
	algoRepo := algo_repo.New(db)
	keyRepo := key_repo.New(db)
	linkRepo := link_repo.New(db)

	numSrv := num_srv.New(numRepo)
	algoFactory := algo_service.NewFactory(algoRepo, numSrv)
	algoSrv, err := algoFactory.Get(*algorithm)
	keySrv := key_service.New(keyRepo, algoSrv)
	linkSrv := link_service.New(linkRepo, keySrv)

	//
	backend := server.NewBackend(linkSrv)

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
