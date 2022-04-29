package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/makhmudvazeez/go-postgres/models"
	"github.com/gorilla/mux"
)

type Response struct {
	Id      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

const UNABLE_TO_CONVERT = "Unable to convert the string into int. %v"
const UNABLE_TO_DECODE = "Unable to decode the requst body. %v"

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully created connection")
	return db
}

func StockIndex(rw http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStock()

	if err != nil {
		log.Fatalf("Unable to get the stock. %v", err)
	}

	json.NewEncoder(rw).Encode(stocks)
}

func StockShow(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf(UNABLE_TO_CONVERT, err)
	}

	stock, err := getStock(int(id))

	if err != nil {
		log.Fatal("Unable to get stock. %v", err)
	}

	json.NewEncoder(rw).Encode(stock)
}

func StockStore(rw http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal(UNABLE_TO_DECODE, err)
	}

	insertId := insertStock(stock)

	response := Response{
		Id: insertId,
		Message: "Successfully created",
	}

	json.NewEncoder(rw).Encode(response)
}

func StockUpdate(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf(UNABLE_TO_CONVERT, err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf(UNABLE_TO_DECODE, err)
	}

	updatedData, _ := updateStock(id, stock)

	msg := fmt.Sprintf("Stock updated successfully. Total rows/records affected %v", updatedData)

	response := Response {
		Id: id,
		Message: msg,
	}

	json.NewEncoder(rw).Encode(response)
}

func StockDestroy(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf(UNABLE_TO_CONVERT, err)
	}

	deletedRows, _ := destroyStock(id)

	msg := fmt.Sprintf("Stock deleted successfully. Total rows/Records %v", deletedRows)

	response := Response {
		Id: id,
		Message: msg,
	}

	json.NewEncoder(rw).Encode(response)

}
