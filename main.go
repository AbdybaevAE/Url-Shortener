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
	"github.com/abdybaevae/url-shortener/pkg/conf"
	"github.com/abdybaevae/url-shortener/pkg/database"
	"github.com/abdybaevae/url-shortener/pkg/migrations"
	"github.com/go-redis/redis/v8"

	algo_repo "github.com/abdybaevae/url-shortener/pkg/repos/algo"
	algo_srv "github.com/abdybaevae/url-shortener/pkg/services/algo"

	key_repo "github.com/abdybaevae/url-shortener/pkg/repos/key"
	key_srv "github.com/abdybaevae/url-shortener/pkg/services/key"

	num_repo "github.com/abdybaevae/url-shortener/pkg/repos/number"
	num_srv "github.com/abdybaevae/url-shortener/pkg/services/number"

	link_repo "github.com/abdybaevae/url-shortener/pkg/repos/link"
	link_srv "github.com/abdybaevae/url-shortener/pkg/services/link"

	cache_srv "github.com/abdybaevae/url-shortener/pkg/services/cache"
	token_srv "github.com/abdybaevae/url-shortener/pkg/services/token"

	user_repo "github.com/abdybaevae/url-shortener/pkg/repos/user"
	user_srv "github.com/abdybaevae/url-shortener/pkg/services/user"

	pbLinks "github.com/abdybaevae/url-shortener/proto/links"
	pbUsers "github.com/abdybaevae/url-shortener/proto/users"
	"github.com/abdybaevae/url-shortener/server"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
	config, err := conf.Load(".")
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.Conn(config)
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
	userRepo := user_repo.New(db)

	tokenSrv := token_srv.New()
	userSrv := user_srv.New(userRepo, tokenSrv)
	numSrv := num_srv.New(numRepo)
	algoFactory := algo_srv.NewFactory(algoRepo, numSrv)
	algoSrv, err := algoFactory.Get(config.Algo)
	if err != nil {
		log.Fatalf("Cannot init config algo %s %v \n", config.Algo, err)
	}
	keySrv := key_srv.New(keyRepo, algoSrv)
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
	})
	cacheSrv := cache_srv.New(redisClient)
	linkSrv := link_srv.New(linkRepo, keySrv, cacheSrv)

	//
	backend := server.NewBackend(linkSrv, userSrv)

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
	pbLinks.RegisterLinkServiceServer(s, backend)
	pbUsers.RegisterUsersServiceServer(s, backend)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
