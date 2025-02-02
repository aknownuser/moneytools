package api

import (
	"encoding/json"
	"moneytool/models"
	"moneytool/storage"
	"net/http"
)

// Handler contains the API handlers and storage reference
type Handler struct {
	store storage.Storage
}

func NewHandler(store storage.Storage) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.store.GetAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(accounts)
}

func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	account, err := h.store.GetAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(account)
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)
	err := h.store.CreateAccount(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(account)
}
