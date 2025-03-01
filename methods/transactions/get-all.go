package transactions

import (
	"encoding/json"
	"net/http"
	"simple-api/db"
	"simple-api/models"
	"simple-api/router"
)

func GetAllTransactions(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if req.URL.Path != "/transactions" {
		router.ErrorJson(res, "Path not found", http.StatusBadRequest)
		return
	}

	conn, errDb := db.Connection()
	if errDb != nil {
		router.ErrorJson(res, "Internal server error connecting to database", http.StatusInternalServerError)
		return
	}

	var transactions []models.Transactions
	query := `SELECT * FROM transactions`
	rows, err := conn.Query(query)
	if err != nil {
		router.ErrorJson(res, "Internal server error querying database", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var transaction models.Transactions
		err := rows.Scan(&transaction.Code, &transaction.Date, &transaction.Quantity, &transaction.TotalPrice, &transaction.Discount, &transaction.Status, &transaction.Payment, &transaction.UserId, &transaction.ProductId)
		if err != nil {
			router.ErrorJson(res, "Internal server error scanning row", http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, transaction)
	}

	response := map[string]interface{}{
		"message": "Success",
		"data":    transactions,
	}

	res.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		router.ErrorJson(res, "Internal server error encoding response", http.StatusInternalServerError)
		return
	}

}
