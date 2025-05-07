package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josseycodes1/ecom-api/types"
	"github.com/josseycodes1/ecom-api/utils"
)

type Handler struct {
	store *types.UserStore
}

func Newhandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	//get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//check if user exist by checking the database

	//create user if it doesnt exist

}
