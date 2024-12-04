package routes

import (
	"github.com/abc/bishal/controllers"
	"github.com/gorilla/mux"
)

func InitAllUserRoutes(router *mux.Router) {
router.HandleFunc("/users",controllers.HandleUsersRequest)
}