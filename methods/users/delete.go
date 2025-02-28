package users

import (
	"database/sql"
	"errors"

	"encoding/json"
	"fmt"
	"net/http"
	"simple-api/db"
	"simple-api/router"

	"github.com/go-sql-driver/mysql"
)

func DeleteUserById(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(req.URL.Path) <= len("/user/delete/") {
		router.ErrorJson(res, "Not found", http.StatusNotFound)
		return
	}

	id := req.URL.Path[len("/user/delete/"):]

	query := "DELETE FROM api1.users WHERE id = ?"

	conn, errConn := db.Connection()
	if errConn != nil {
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

	result, errExc := conn.Exec(query, id)
	if errExc != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(errExc, &mysqlErr) {
			if mysqlErr.Number == 1452 {
				router.ErrorJson(res, "Foreign key constraint fails", http.StatusBadRequest)
				return
			}
			if mysqlErr.Number == 1451 {
				router.ErrorJson(res, "Foreign key constraint fails", http.StatusBadRequest)
				return
			}
		}

		router.ErrorJson(res, "Internal server error execute", http.StatusInternalServerError)
		fmt.Fprintf(res, "Error: %v", errExc)
		return
	}

	affected, errAffected := result.RowsAffected()
	if errAffected != nil {
		router.ErrorJson(res, "Internal server error affected", http.StatusInternalServerError)
		return
	}

	if affected == 0 {
		var exist bool

		err := conn.QueryRow("SELECT EXISTS(SELECT * FROM api1.users WHERE id = ?)", id).Scan(&exist)
		if err != nil {
			router.ErrorJson(res, "Internal server error exist", http.StatusInternalServerError)
			return
		}

		if !exist {
			router.ErrorJson(res, "User not found", http.StatusNotFound)
			return
		}

		router.ErrorJson(res, "No rows affected", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"message": "Success deleted user",
	}

	res.Header().Set("Content-Type", "application/json")
	errJson := json.NewEncoder(res).Encode(response)
	if errJson != nil {
		router.ErrorJson(res, "Internal server error json", http.StatusInternalServerError)
		return
	}

}
