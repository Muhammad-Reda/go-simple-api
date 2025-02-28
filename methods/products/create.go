package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-api/db"
	"simple-api/models"
	"simple-api/router"
	"strconv"
)

func CreateNewProduct(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if req.URL.Path != "/product/add" {
		router.ErrorJson(res, "Path not found", http.StatusBadRequest)
		return
	}

	errParse := req.ParseForm()
	if errParse != nil {
		router.ErrorJson(res, "Can not parse form data", http.StatusBadRequest)
		return
	}

	formData := req.PostForm
	priceStr := formData.Get("price")
	price, err := strconv.ParseInt(priceStr, 10, 64)
	if err != nil {
		router.ErrorJson(res, "Invalid price format", http.StatusBadRequest)
		return
	}

	var product = models.Product{
		Code:     formData.Get("code"),
		Name:     formData.Get("name"),
		Category: formData.Get("category"),
		Price:    price,
	}

	conn, errDb := db.Connection()
	if errDb != nil {
		router.ErrorJson(res, "Internal server error connecting database", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	query := "INSERT INTO products (code, name, category, price) VALUES (?, ?, ?, ?);"
	_, errExc := conn.Exec(query, product.Code, product.Name, product.Category, product.Price)
	if errExc != nil {
		router.ErrorJson(res, "Internal server error executing query", http.StatusInternalServerError)
		fmt.Fprintf(res, "Error: %v", errExc)
		return
	}

	response := map[string]string{
		"message": "Success",
	}

	res.Header().Set("Content-Type", "application/json")
	if errJson := json.NewEncoder(res).Encode(response); errJson != nil {
		router.ErrorJson(res, "Internal server error encoding json", http.StatusInternalServerError)
		return
	}

}
