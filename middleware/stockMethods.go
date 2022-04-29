package middleware

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/makhmudvazeez/go-postgres/models"
)

func getAllStock() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()

	stocks := []models.Stock{}

	rows, err := db.Query(`SELECT * FROM stocks`)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		stock := models.Stock{}
		err := rows.Scan(&stock.Id, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func getStock(id int) (models.Stock, error) {
	db := createConnection()
	defer db.Close()

	stock := models.Stock{}
	err := db.QueryRow(`SELECT * FROM stocks WHERE id=$1`, id).Scan(&stock.Id, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No row were returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return stock, nil
}

func insertStock(stock models.Stock) (id int){
	db := createConnection()
	defer db.Close()

	err := db.QueryRow(`INSERT INTO stock(name, price, company) VALUSE($1, $2 ,$3) RETURNING id`, stock.Name, stock.Price, stock.Company).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id

}

func updateStock(id int, stock models.Stock) (int, error){
	db := createConnection()
	defer db.Close()

	row, err := db.Exec(`UPDATE stocks SET name=$2, price=$3, company=$4 WHERE id=$1`, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatal("Unable to execute the query. %v", err)
	}

	rowsAffected, err := row.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/records has been affected %v", rowsAffected)
	return 1, nil
}

func destroyStock(id int) (int, error){
	db := createConnection()
	defer db.Close()

	row, err := db.Exec(`DELETE FROM stocks WHERE id=$1`, id)

	if err != nil {
		log.Fatal("Unable to execute the query. %v", err)
	}

	rowsAffected, err := row.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/records has been affected %v", rowsAffected)
	return 1, nil
}
