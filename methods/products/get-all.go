package products

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"simple-api/db"
	"simple-api/models"
	"simple-api/router"
)

func GetAllProducts(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if req.URL.Path != "/products" {
		router.ErrorJson(res, "Not found", http.StatusNotFound)
		return
	}

	var products []models.Product
	query := "SELECT * FROM api1.products"

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

	rows, errQuery := conn.Query(query)
	if errQuery != nil {
		router.ErrorJson(res, "Internal server error Query", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Code, &product.Name, &product.Category, &product.Price, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			router.ErrorJson(res, "Internal server error scan", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	response := map[string]interface{}{
		"message": "Success",
		"data":    products,
	}

	res.Header().Set("Content-Type", "application/json")

	errJson := json.NewEncoder(res).Encode(response)
	if errJson != nil {
		router.ErrorJson(res, "Internal server error json", http.StatusInternalServerError)
		return
	}

}
