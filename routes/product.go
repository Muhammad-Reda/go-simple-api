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
	router.HandleFunc("/product/add", cors.IPFilter(products.CreateNewProduct, rtr.ListPattern))
	router.HandleFunc("/product/update/{id}", cors.IPFilter(products.UpdateProduct, rtr.ListPattern))
	router.HandleFunc("/product/delete/{id}", cors.IPFilter(products.DeleteProduct, rtr.ListPattern))
}
