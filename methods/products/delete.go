package products

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"simple-api/db"
	"strconv"
)

func DeleteProduct(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if len(req.URL.Path) <= len("/product/delete/") {
		http.Error(res, "Product ID is required", http.StatusBadRequest)
		return
	}

	id := req.URL.Path[len("/product/delete/"):]

	_, errId := strconv.Atoi(id)
	if errId != nil {
		http.Error(res, "Invalid product ID, ID must be int", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM products WHERE id = ?"
	conn, errDb := db.Connection()
	if errDb != nil {
		http.Error(res, "Internal Server error connecting database", http.StatusInternalServerError)
		return
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			http.Error(res, "Internal Server error closing database connection", http.StatusInternalServerError)
			return
		}
	}(conn)

	result, errExc := conn.Exec(query, id)
	if errExc != nil {
		http.Error(res, "Internal Server error executing query", http.StatusInternalServerError)
		return
	}

	rowsAffected, errAffected := result.RowsAffected()
	if errAffected != nil {
		http.Error(res, "Internal Server error getting affected rows", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"message": "Product deleted successfully"}

	res.Header().Set("Content-Type", "application/json")
	errJson := json.NewEncoder(res).Encode(response)
	if errJson != nil {
		http.Error(res, "Internal Server error encoding response", http.StatusInternalServerError)
		return
	}

}
