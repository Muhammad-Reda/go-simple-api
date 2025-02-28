package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-api/cors"
	"simple-api/router"
	rtr "simple-api/router"
	"simple-api/routes"
)

func main() {
	mainRouter := rtr.MainRouter

	mainRouter.HandleFunc("/", cors.IPFilter(func(res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			router.ErrorJson(res, "Not found", http.StatusNotFound)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "Hello world",
		}
		err := json.NewEncoder(res).Encode(response)
		if err != nil {
			rtr.ErrorJson(res, "Internal server error json", http.StatusInternalServerError)
			return
		}
	}, rtr.ListPattern))

	routes.UserRoutes()
	routes.ProductRoutes()

	err := http.ListenAndServe(":8080", rtr.MainRouter)
	if err != nil {
		fmt.Println("Can not serve server")
		return
	}
}
