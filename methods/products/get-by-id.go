package products

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"simple-api/db"
	"simple-api/models"
	"simple-api/router"
)

func GetProductByCode(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(req.URL.Path) <= len("/product/") {
		router.ErrorJson(res, "Path not found", http.StatusNotFound)
		return
	}

	id := req.URL.Path[len("/product/"):]
	var product models.Product

	query := "SELECT * FROM products WHERE id = ?"

	conn, errDb := db.Connection()
	if errDb != nil {
		router.ErrorJson(res, "Internal server error connecting database", http.StatusInternalServerError)
	}
	defer conn.Close()

	errQuery := conn.QueryRow(query, id).Scan(&product.Id, &product.Code, &product.Name, &product.Category, &product.Price, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)

	if errQuery != nil {
		if errors.Is(errQuery, sql.ErrNoRows) {
			router.ErrorJson(res, "Product not found", http.StatusNotFound)
			return
		}

		router.ErrorJson(res, "Internal server error Scan", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Success",
		"product": product,
	}

	errJson := json.NewEncoder(res).Encode(response)
	if errJson != nil {
		router.ErrorJson(res, "Internal server error encoding json", http.StatusInternalServerError)
		return
	}
}
