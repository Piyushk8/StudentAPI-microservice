package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/piyushk8/StudentAPI/internal/Storage/sqlite"
	"github.com/piyushk8/StudentAPI/internal/config"
	student "github.com/piyushk8/StudentAPI/internal/http/handlers"
)

func test() {
	messageChain := make(chan string)

	messageChain <- "ping"

	msg := <-messageChain

	fmt.Println(msg)
}

func main() {
	// test()
	// // loaidng configs
	cfg := config.MUSTLoad()
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))



	router := http.NewServeMux()

	router.HandleFunc("GET /api/student", student.New(storage))

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server:", err)
		}
	}()

	slog.Info("server started", slog.String("addr", cfg.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done // wait for Ctrl+C

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	slog.Info("shutting down server")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", slog.String("error", err.Error()))
	} else {
		slog.Info("server exited gracefully")
	}
}
