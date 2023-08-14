package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	infraDatabase "github.com/jeanSagaz/server/internal/infra/database"
	"github.com/jeanSagaz/server/internal/webserver/handlers"
	"github.com/jeanSagaz/server/pkg/database"
)

var db database.Database

func init() {
	db.AutoMigrateDb = true
}

func main() {
	dbConnection, err := db.Connect()
	if err != nil {
		log.Fatalf("error connecting to DB")
	}

	sqlDB, err := dbConnection.DB()
	if err != nil {
		log.Fatalf("error sql to DB")
	}
	defer sqlDB.Close()

	usdBrlRateRepository := infraDatabase.NewUsdBrlRepositoryDb(dbConnection)
	exchangeRateHandler := handlers.NewExchangeRateHandler(usdBrlRateRepository)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/cotacao", exchangeRateHandler.GetQuotationTax)
	log.Fatal(http.ListenAndServe(":8080", router))
}
