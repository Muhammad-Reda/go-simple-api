package routes

import (
	"simple-api/cors"
	"simple-api/methods/products"
	rtr "simple-api/router"
)

func ProductRoutes() {
	router := rtr.MainRouter

	router.HandleFunc("/products", cors.IPFilter(products.GetAllProducts, rtr.ListPattern))
	router.HandleFunc("/product/{id}", cors.IPFilter(products.GetProductByCode, rtr.ListPattern))
}
