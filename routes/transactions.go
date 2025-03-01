package routes

import (
	"simple-api/cors"
	"simple-api/methods/transactions"
	rtr "simple-api/router"
)

func TransactionsRoutes() {
	router := rtr.MainRouter

	router.HandleFunc("/transactions", cors.IPFilter(transactions.GetAllTransactions, rtr.ListPattern))
	router.HandleFunc("transaction/{id}", cors.IPFilter(transactions.GetTransactionById, rtr.ListPattern))
	router.HandleFunc("/transaction/add", cors.IPFilter(transactions.CreateTransaction, rtr.ListPattern))
}
