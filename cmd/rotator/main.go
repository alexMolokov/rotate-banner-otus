package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexMolokov/rotate-banner-otus/internal/app/rotator"
	"github.com/alexMolokov/rotate-banner-otus/internal/app/rotator/observer"
	configApp "github.com/alexMolokov/rotate-banner-otus/internal/config"
	"github.com/alexMolokov/rotate-banner-otus/internal/logger"
	internalgrpc "github.com/alexMolokov/rotate-banner-otus/internal/server/grpc"
	rs "github.com/alexMolokov/rotate-banner-otus/internal/storage/rotator"
)

var configFile string

func init() {
	flag.StringVar(
		&configFile,
		"config",
		"./configs/rotator.json",
		"Path to configuration file",
	)
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	cfg, err := configApp.NewRotatorConfig(configFile)
	if err != nil {
		fmt.Printf("Can't load config: %v", err)
		os.Exit(1)
	}

	lgr, err := logger.New(&cfg.Logger)
	if err != nil {
		fmt.Printf("Can't create lgr: %v", err)
		os.Exit(1)
	}

	storage := rs.NewRotatorStorage(cfg.DB)
	err = storage.Connect()
	if err != nil {
		log.Fatalf("Can't create pool connect to storage: %v", err)
	}
	defer func() {
		err = storage.Close()
		if err != nil {
			lgr.Error("failed to close pool connection to storage: " + err.Error())
		}
	}()

	handler := observer.NewAMQPObserver(cfg.Queue, lgr)
	err = handler.Init()
	if err != nil {
		lgr.Error("Can't init connect rabbit: %v", err)
	}
	defer handler.Close()

	app := approtator.NewAppRotator(lgr, storage)
	app.AddHandler(handler)

	tcpAddr := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)
	grpcServer := internalgrpc.NewServer(lgr, app, tcpAddr)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		lgr.Info("Service GRPC bannerRotator is running...")

		if err := grpcServer.Start(); err != nil {
			lgr.Error("failed to start GRPC server: " + err.Error())
			cancel()
		}
	}()

	<-ctx.Done()

	fmt.Println("Graceful shutdown start")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := grpcServer.Stop(ctx); err != nil {
		lgr.Error("failed to stop GRPC bannerRotator service: " + err.Error())
	} else {
		msg := "Service GRPC bannerRotator is stopped"
		fmt.Println(msg)
		lgr.Info(msg)
	}
}
