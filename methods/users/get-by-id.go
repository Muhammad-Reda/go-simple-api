package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"simple-api/db"
	"simple-api/models"
	"simple-api/router"
)

func GetUserById(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}

	if len(req.URL.Path) <= len("/user/") {
		router.ErrorJson(res, "Not Found", http.StatusNotFound)
		return

	}

	id := req.URL.Path[len("/user/"):]

	query := "SELECT * FROM api1.users WHERE id = ?"

	conn, errorDb := db.Connection()
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			router.ErrorJson(res, "Internal server error closing db", http.StatusInternalServerError)
			return
		}
	}(conn)

	if errorDb != nil {
		router.ErrorJson(res, "Internal server error connecting db", http.StatusInternalServerError)
		return

	}

	var user models.User

	errScan := conn.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Address, &user.Telephone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if errScan != nil {
		if errors.Is(errScan, sql.ErrNoRows) {
			router.ErrorJson(res, "User Not found", http.StatusNotFound)
			return
		}

		router.ErrorJson(res, "Internal server error scaning", http.StatusInternalServerError)
		fmt.Fprintf(res, "Eror: %s", errScan)
		return
	}

	response := map[string]interface{}{
		"message": "Success",
		"data":    user,
	}

	res.Header().Set("Content-Type", "application/json")

	errJson := json.NewEncoder(res).Encode(response)
	if errJson != nil {
		router.ErrorJson(res, "Internal server error json", http.StatusInternalServerError)
		return
	}
}
