package transactions

import (
	"encoding/json"
	"net/http"
	"simple-api/db"
	"simple-api/models"
	"simple-api/router"
)

func GetTransactionById(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if len(req.URL.Path) <= len("/transaction/") {
		router.ErrorJson(res, "Path not found and code required", http.StatusBadRequest)
		return
	}

	conn, errDb := db.Connection()
	if errDb != nil {
		router.ErrorJson(res, "Internal server error connecting to database", http.StatusInternalServerError)
		return
	}

	transactionCode := req.URL.Path[len("/transaction/")]

	var transaction models.Transactions
	query := `SELECT * FROM api1.transactions WHERE code = ?`

	row := conn.QueryRow(query, transactionCode)
	errScan := row.Scan(&transaction.Code, &transaction.Date, &transaction.Quantity, &transaction.TotalPrice, &transaction.Discount, &transaction.Status, &transaction.Payment, &transaction.UserId, &transaction.ProductId)
	if errScan != nil {
		router.ErrorJson(res, "Transaction not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"message": "Success",
		"data":    transaction,
	}

	res.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(res).Encode(response)
	if err != nil {
		router.ErrorJson(res, "Internal server error encoding response", http.StatusInternalServerError)
		return
	}
}
