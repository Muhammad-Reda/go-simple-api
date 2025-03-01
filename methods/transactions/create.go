package transactions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"simple-api/db"
	"simple-api/models"
	"simple-api/router"
	"strconv"
	"time"
)

func CreateTransaction(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if req.URL.Path != "/transaction/add" {
		router.ErrorJson(res, "Path not found", http.StatusBadRequest)
		return
	}

	var transaction models.Transactions
	errParse := req.ParseForm()
	if errParse != nil {
		router.ErrorJson(res, "Error parsing form data", http.StatusBadRequest)
		return
	}
	formData := req.Form

	transaction.Code = formData.Get("code")

	strDate := formData.Get("date")
	if strDate == "" {
		router.ErrorJson(res, "Date is required", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		router.ErrorJson(res, "Invalid date format", http.StatusBadRequest)
		fmt.Fprintf(res, "strdate: %s", strDate)
		return
	}
	transaction.Date = date

	strQuantity := formData.Get("quantity")
	quantity, errQuantity := strconv.Atoi(strQuantity)
	if errQuantity != nil {
		router.ErrorJson(res, "Invalid quantity format", http.StatusBadRequest)
		return
	}
	transaction.Quantity = quantity

	strDiscount := formData.Get("discount")
	discount, errDiscount := strconv.Atoi(strDiscount)
	if errDiscount != nil {
		router.ErrorJson(res, "Invalid discount format", http.StatusBadRequest)
		return
	}
	transaction.Discount = discount

	transaction.Status = formData.Get("status")
	transaction.Payment = formData.Get("payment")

	strUserId := formData.Get("user_id")
	userId, errUserId := strconv.ParseInt(strUserId, 10, 64)
	if errUserId != nil {
		router.ErrorJson(res, "Invalid user ID format", http.StatusBadRequest)
		return
	}
	transaction.UserId = userId

	strProductId := formData.Get("product_id")
	productId, errProductId := strconv.Atoi(strProductId)
	if errProductId != nil {
		router.ErrorJson(res, "Invalid product ID format", http.StatusBadRequest)
		return
	}
	transaction.ProductId = productId

	conn, errDb := db.Connection()
	if errDb != nil {
		router.ErrorJson(res, "Internal server error connecting to database", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var price int
	priceQuery := `SELECT price FROM api1.products WHERE id = ?`
	priceRow := conn.QueryRow(priceQuery, productId)
	errScan := priceRow.Scan(&price)
	if errScan != nil {
		router.ErrorJson(res, "Product not found", http.StatusNotFound)
		return
	}

	query := `INSERT INTO api1.transactions (code, date, quantity, total_price, discount, status, payment, user_id, product_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, errExec := conn.Exec(query, transaction.Code, transaction.Date, transaction.Quantity, price, transaction.Discount, transaction.Status, transaction.Payment, transaction.UserId, transaction.ProductId)
	if errExec != nil {
		var mySqlError *mysql.MySQLError

		if errors.As(errExec, &mySqlError) {
			if mySqlError.Number == 1062 {
				router.ErrorJson(res, "Duplicate entry", http.StatusBadRequest)
				return
			}

		}

		router.ErrorJson(res, "Internal server error inserting data", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Transaction created successfully",
	}

	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		router.ErrorJson(res, "Internal server error encoding response", http.StatusInternalServerError)
		return
	}
}
