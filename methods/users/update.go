package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"simple-api/db"
	"simple-api/router"
	"strings"
	"time"
)

func UpdateUserById(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPatch {
		router.ErrorJson(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(req.URL.Path) <= len("/user/update/") {
		router.ErrorJson(res, "Not found", http.StatusNotFound)
		return
	}

	id := req.URL.Path[len("/user/update/"):]

	errParseForm := req.ParseForm()
	if errParseForm != nil {
		router.ErrorJson(res, "Bad form data", http.StatusBadRequest)
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

	var setValues []string
	var args []interface{}

	for key, val := range user {
		if val == "" {
			continue
		}
		setValues = append(setValues, key+" = ?")
		args = append(args, val)
	}

	if len(setValues) <= 0 {
		router.ErrorJson(res, "No req form data", http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	args = append(args, currentTime)
	args = append(args, id)

	query := "UPDATE api1.users SET " + strings.Join(setValues, ",") + ", updated_at = ? WHERE id = ?"

	conn, errDb := db.Connection()
	if errDb != nil {
		router.ErrorJson(res, "Internal server error connecting to database", http.StatusInternalServerError)
		return
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			router.ErrorJson(res, "Internal server error closing database", http.StatusInternalServerError)
			return
		}
	}(conn)

	result, errExc := conn.Exec(query, args...)
	if errExc != nil {
		router.ErrorJson(res, "Internal server error execute", http.StatusInternalServerError)
		return
	}

	affected, errAffected := result.RowsAffected()
	if errAffected != nil {
		router.ErrorJson(res, "Internal server error Affected", http.StatusInternalServerError)
		return
	}

	if affected == 0 {
		var exist bool

		errScan := conn.QueryRow("SELECT EXISTS(SELECT * FROM api1.users WHERE id = ?)", id).Scan(&exist)
		if errScan != nil {
			router.ErrorJson(res, "Internal server error Scan", http.StatusInternalServerError)
			return
		}

		if !exist {
			router.ErrorJson(res, "User not found", http.StatusNotFound)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "updated success but no rows affected",
		}
		err := json.NewEncoder(res).Encode(response)
		if err != nil {
			router.ErrorJson(res, "Internal server error send json", http.StatusInternalServerError)
			return
		}

		return

	}

	req.Header.Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "Success updates user",
	}
	err := json.NewEncoder(res).Encode(response)
	if err != nil {
		router.ErrorJson(res, "Internal server error send json", http.StatusInternalServerError)
		return
	}

}
