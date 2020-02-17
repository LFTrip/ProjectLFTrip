package controllers

import (
	"net/http"

	"github.com/LFTrip/ProjectLFTrip/api/responses"
)

// Home : server response
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
