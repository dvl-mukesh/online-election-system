package router

import (
	"net/http"
	"online-election-system/controller"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/api/add-user", controller.AddUser)
	router.HandleFunc("/api/verify-user", controller.VerifyUser)
	router.HandleFunc("/api/search-one-user", controller.SearchOneUser)
	router.HandleFunc("/api/search-multiple-user", controller.SearchMultipleUser)
	router.HandleFunc("/api/delete-user", controller.DeleteUser)

	return router
}
