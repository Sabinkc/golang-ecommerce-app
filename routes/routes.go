package routes

import "github.com/gorilla/mux"

func InitALLRoutes()*mux.Router {
	newRouter := mux.NewRouter()
	InitAllPracticRoutes(newRouter)
	InitAllUserRoutes(newRouter)
	return newRouter
}