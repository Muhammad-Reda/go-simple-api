package products

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"simple-api/db"
	"simple-api/router"
	"strconv"
	"strings"
	"time"
)

func UpdateProduct(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPatch {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(req.URL.Path) <= len("/product/update/") {
		router.ErrorJson(res, "Path not found", http.StatusNotFound)
		return
	}

	errParse := req.ParseForm()
	if errParse != nil {
		router.ErrorJson(res, "Can not parse form data", http.StatusBadRequest)
		return
	}

	id := req.URL.Path[len("/product/update/"):]
	formData := req.Form

	priceStr := formData.Get("price")

	price, errParseInt := strconv.ParseInt(priceStr, 10, 64)
	if errParseInt != nil {
		router.ErrorJson(res, "Invalid price format", http.StatusBadRequest)
		return
	}

	var product = map[string]interface{}{
		"code":     formData.Get("code"),
		"name":     formData.Get("name"),
		"category": formData.Get("category"),
		"price":    price,
	}

	var setValues []string
	var args []interface{}

	for key, value := range product {
		if value == "" {
			continue
		}
		setValues = append(setValues, key+" = ?")
		args = append(args, value)
	}

	if len(setValues) <= 0 {
		router.ErrorJson(res, "No req form data", http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	args = append(args, currentTime)
	args = append(args, id)

	query := "UPDATE api1.products SET " + strings.Join(setValues, ",") + ", updated_at = ? WHERE id = ?"

	conn, errDb := db.Connection()
	if errDb != nil {
		router.ErrorJson(res, "Internal server error connecting database", http.StatusInternalServerError)
		return
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			router.ErrorJson(res, "Internal server error closing database", http.StatusInternalServerError)
			return
		}
	}(conn)

	result, errExec := conn.Exec(query, args...)

	if errExec != nil {
		router.ErrorJson(res, "Internal server error executing query", http.StatusInternalServerError)
		return
	}
	rowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		router.ErrorJson(res, "Internal server error getting rows affected", http.StatusInternalServerError)
		return
	}

	if rowsAffected <= 0 {
		var exist bool
		query := "SELECT EXISTS(SELECT 1 FROM products WHERE id = ?)"
		errScan := conn.QueryRow(query, id).Scan(&exist)
		if errScan != nil {
			router.ErrorJson(res, "Internal server error scanning", http.StatusInternalServerError)
			return
		}
		if !exist {
			router.ErrorJson(res, "Product not found", http.StatusNotFound)
			return
		}

		response := map[string]string{
			"message": "Updated success but no data changed",
		}
		res.Header().Set("Content-Type", "application/json")
		errJson := json.NewEncoder(res).Encode(response)
		if errJson != nil {
			router.ErrorJson(res, "Internal server error encoding json", http.StatusInternalServerError)
			return
		}

		return
	}

	response := map[string]string{
		"message": "Success",
	}

	res.Header().Set("Content-Type", "application/json")
	errJson := json.NewEncoder(res).Encode(response)
	if errJson != nil {
		router.ErrorJson(res, "Internal server error encoding json", http.StatusInternalServerError)
		return
	}

}
