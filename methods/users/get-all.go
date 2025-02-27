package users

import (
	"encoding/json"
	"net/http"
	"simple-api/db"
	"simple-api/models"
	"simple-api/router"
)

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if req.URL.Path != "/users" {
		router.ErrorJson(res, "Not found", http.StatusNotFound)
		return

	}

	conn, errDb := db.Connection()
	defer conn.Close()
	if errDb != nil {
		router.ErrorJson(res, "Internal server error", http.StatusInternalServerError)
		return

	}

	query := "SELECT * FROM users"
	var users []models.User

	rows, errQuery := conn.Query(query)
	if errQuery != nil {
		router.ErrorJson(res, "Internal server error", http.StatusInternalServerError)
		return

	}

	for rows.Next() {
		var user models.User

		rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Address, &user.Telephone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

		users = append(users, user)
	}

	response := map[string]interface{}{
		"message": "Success",
		"data":    users,
	}

	res.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(res).Encode(response)
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

}
