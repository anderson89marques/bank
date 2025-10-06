package rest

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

	"github.com/anderson89marques/bank/config"
	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) {
	err := config.ParseEnv()
	if err != nil {
		log.Fatal("error to load environment variables", err)
	}

	engine := gin.New()
	RegisterRoutes(engine)

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", config.GetEnv().AppHost, config.GetEnv().AppPort),
		Handler:           engine,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout de 5 seconds
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can`t be catch, so don`t need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server Shutdown", slog.String("error", err.Error()))
	}

	// catching ctx.Done(). timeout of 5 seconds
	select {
	case <-ctx.Done():
		slog.Info("Timeout of 5 seconds.")
	}
	slog.Info("Server Shutdown")
}
