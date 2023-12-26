package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/UnitedIngvar/onmi_test/internal/config"
	"github.com/UnitedIngvar/onmi_test/internal/lib/slogext"
	"github.com/UnitedIngvar/onmi_test/internal/services/client"
	"github.com/UnitedIngvar/onmi_test/internal/services/extservice"
	"github.com/UnitedIngvar/onmi_test/internal/transport/rest/handlers/sendRequest"
	"github.com/UnitedIngvar/onmi_test/internal/transport/rest/middleware/loggermw"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func StartServer(log *slog.Logger, cfg *config.Config) {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(loggermw.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(cfg.Timeout))

	client := client.NewClient(extservice.NewMyMockService())
	router.Get("/docs/*", httpSwagger.WrapHandler.ServeHTTP)
	router.Post("/send-request", sendRequest.NewHandler(log, client))

	log.Info("starting server", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server")
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	// TODO: move timeout to config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", slogext.Error(err))

		return
	}

	log.Info("server stopped")
}
