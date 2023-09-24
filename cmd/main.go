package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aerostatas/interaction-go-kata/docs"
	"github.com/aerostatas/interaction-go-kata/internal/api/handler"
	"github.com/aerostatas/interaction-go-kata/internal/repository"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if err := NewAPICmd().Execute(); err != nil {
		log.Fatal(err)
	}
}

type apiOptions struct {
	address    string
	sqlitePath string
	docs       bool
}

func NewAPICmd() *cobra.Command {
	opts := &apiOptions{}

	cmd := &cobra.Command{
		Use:   "server",
		Short: "REST API endpoint that allows users to create events",
		RunE: func(c *cobra.Command, args []string) error {
			db, err := prepareDB(opts.sqlitePath)
			if err != nil {
				return fmt.Errorf("prepare DB: %w", err)
			}

			repositoryFactory := repository.NewRepositoryFactory(db)
			handlerFactory := handler.NewHandlerFactory(repositoryFactory)

			server, err := prepareHTTPServer(opts.address, opts.docs, handlerFactory)
			if err != nil {
				return fmt.Errorf("prepare HTTP server: %w", err)
			}

			shutdown := make(chan error)
			osSignals := make(chan os.Signal, 1)
			signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
			go func() {
				s := <-osSignals
				log.Printf("shutdown signal received: %s\n", s)
				shutdown <- nil
			}()

			go func() {
				if err := server.ListenAndServe(); err != http.ErrServerClosed {
					shutdown <- err
				}
			}()

			if err := <-shutdown; err != nil {
				log.Printf("shutdown due to error: %v\n", err)
			}

			return server.Shutdown(context.Background())
		},
	}

	cmd.Flags().StringVar(&opts.address, "address", "localhost:8080", "address to launch the HTTP server on")
	cmd.Flags().StringVar(&opts.sqlitePath, "sqlite-path", "app.sqlite", "path to SQLite DB")
	cmd.Flags().BoolVar(&opts.docs, "docs", true, "include documentation routes")

	return cmd
}

func prepareDB(sqlitePath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection: %s", err)
	}

	// migrate DB schema
	if err := repository.Migrate(db); err != nil {
		return nil, err
	}

	// seed default data
	if err := repository.Seed(db); err != nil {
		return nil, err
	}

	return db, err
}

// @title Event API
// @version 1.0
// @description REST API that allows users to create events

// @host localhost:8080
// @BasePath /api

// @accept json
func prepareHTTPServer(
	address string,
	includeDocs bool,
	handlerFactory *handler.HandlerFactory,
) (*http.Server, error) {
	eventHandler, err := handlerFactory.EventHandler()
	if err != nil {
		return nil, fmt.Errorf("make event handler: %w", err)
	}

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.Handle("/events", eventHandler.CreateEvent()).Methods("POST")

	if includeDocs {
		if address != "" {
			docs.SwaggerInfo.Host = address
			docs.SwaggerInfo.Schemes = []string{"http"}
		}
		router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	}

	return &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  30 * time.Second,
	}, nil
}
