package router

import (
	"github.com/gorilla/mux"
	"github.com/makhmudvazeez/go-postgres/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stocks", middleware.StockIndex).Methods("GET")
	router.HandleFunc("/api/stocks{id}", middleware.StockShow).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stocks", middleware.StockStore).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stocks", middleware.StockUpdate).Methods("PUT", "PATCH", "OPTIONS")
	router.HandleFunc("/api/stocks{id}", middleware.StockDestroy).Methods("DELETE", "OPTIONS")

	return router
}
