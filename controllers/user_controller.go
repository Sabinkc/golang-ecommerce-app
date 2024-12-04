package controllers

import (
	"net/http"

	"github.com/abc/bishal/models"
	"github.com/abc/bishal/responses"
)

func HandleUsersRequest(w http.ResponseWriter, r *http.Request) {
	allModels := models.FetAllUsers()
	responses.SuccessResponse(w, allModels, "Data fetched successfully!")

}
