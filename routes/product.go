package routes

import (
	"simple-api/cors"
	"simple-api/methods/products"
	rtr "simple-api/router"
)

func ProductRoutes() {
	router := rtr.MainRouter

	router.HandleFunc("/products", cors.IPFilter(products.GetAllProducts, rtr.ListPattern))
}
