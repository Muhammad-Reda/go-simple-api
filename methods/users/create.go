package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"simple-api/db"
	"simple-api/router"
)

func CreateNewUser(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if req.URL.Path != "/user/add" {
		router.ErrorJson(res, "Not found", http.StatusNotFound)
		return
	}

	err := req.ParseForm()
	if err != nil {
		router.ErrorJson(res, "Can not parse request body", http.StatusBadRequest)
		return
	}

	formData := req.Form

	user := map[string]string{
		"email":     formData.Get("email"),
		"username":  formData.Get("username"),
		"password":  formData.Get("password"),
		"address":   formData.Get("address"),
		"telephone": formData.Get("telephone"),
	}

	for key, value := range user {
		if value == "" {
			router.ErrorJson(res, fmt.Sprintf("%s can't be empty", key), http.StatusBadRequest)
			return
		}
	}

	query := "INSERT INTO api1.users (email, username, password, address, telephone) VALUES (?, ?, ?, ?, ?)"

	conn, errorDb := db.Connection()
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			router.ErrorJson(res, "Internal server error closing database", http.StatusInternalServerError)
			return
		}
	}(conn)

	if errorDb != nil {
		router.ErrorJson(res, "internal server error connecting database", http.StatusInternalServerError)
		return
	}

	_, errExec := conn.Exec(query, user["email"], user["username"], user["password"], user["address"], user["telephone"])
	if errExec != nil {
		var mysqlErr *mysql.MySQLError

		if errors.As(errExec, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				router.ErrorJson(res, "Email or username already exist", http.StatusBadRequest)
			default:
				router.ErrorJson(res, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		router.ErrorJson(res, "Internal server error errexc", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Success",
	}

	errJson := json.NewEncoder(res).Encode(response)
	if errJson != nil {
		router.ErrorJson(res, "Internal server error to send json", http.StatusInternalServerError)
		return
	}

}
