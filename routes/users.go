package routes

import (
	"simple-api/cors"
	"simple-api/methods/users"
	rtr "simple-api/router"
)

func UserRoutes() {
	router := rtr.MainRouter
	router.HandleFunc("/users", cors.IPFilter(users.GetAllUsers, rtr.ListPattern))
	router.HandleFunc("/user/{id}", cors.IPFilter(users.GetUserById, rtr.ListPattern))
	router.HandleFunc("/user/add", cors.IPFilter(users.CreateNewUser, rtr.ListPattern))
	router.HandleFunc("/user/update/{id}", cors.IPFilter(users.UpdateUserById, rtr.ListPattern))
	router.HandleFunc("/user/delete/{id}", cors.IPFilter(users.DeleteUserById, rtr.ListPattern))
}
