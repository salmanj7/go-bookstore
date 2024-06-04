package routes

import (
	"github.com/gorilla/mux"
	"github.com/salmanj7/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.Handlefunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.Getbook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}