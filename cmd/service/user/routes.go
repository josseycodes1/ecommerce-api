package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josseycodes1/ecom-api/types"
)

type Handler struct {
}

func Newhandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}

//get json payload
var payload types.RegisterUserPayload
if err := utils.ParseJSON(r.Body, payload); err != nil {
	utils.WriteError (w, http.StatusBadRequest, err)
}

//check if user exist
//create user if it doesnt exist