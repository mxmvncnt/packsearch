package main

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mxmvncnt/packsearch/server/config"
	"github.com/mxmvncnt/packsearch/server/database"
	"github.com/mxmvncnt/packsearch/server/middleware"
	"github.com/mxmvncnt/packsearch/server/routes"
	"github.com/mxmvncnt/packsearch/server/utils/logger"
	"github.com/rs/cors"
)

func main() {
	logger.Infof("Loaded config '%s'", config.ConfigName)

	router := http.NewServeMux()

	dbPool, dbErr := pgxpool.New(context.Background(), config.DatabaseURL)
	if dbErr != nil {
		logger.Fatalf("Failed to initialize database: %v", dbErr)
	}
	defer dbPool.Close()
	queries := database.New(dbPool)
	dbErr = dbPool.Ping(context.Background())
	if dbErr != nil {
		logger.Fatalf("Failed to connect to database: %v", dbErr)
	}

	routes := routes.NewRoutesHandler(queries)
	router.HandleFunc("GET /packages/list", middleware.Combined(routes.GetPackagesList))
	router.HandleFunc("GET /packages/{id}", middleware.Combined(routes.GetPackage))
	router.HandleFunc("GET /packages/{id}/variations", middleware.Combined(routes.GetVariations))
	router.HandleFunc("GET /search/{search_term}", middleware.Combined(routes.SearchPackage))

	logger.Info("Server started on http://" + config.ServerHostname + ":" + config.ServerPort)
	handler := cors.AllowAll().Handler(router)
	http.ListenAndServe(config.ServerHostname+":"+config.ServerPort, handler)
}
