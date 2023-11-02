package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/1shubham7/go-mongo-crud/handlers"
)

func main() {

	router := httprouter.New()
	router.GET("/users", handlers.GetAllUsers)
	router.GET("/user/:id", handlers.GetUserById)

	router.DELETE("delete/user/:id, handlers.DeleteUserById")

	router.POST("add/user/:id", handlers.AddUserById)

	router.PUT("update/user/:id", handlers.UpdateUserBuId)
	

}
