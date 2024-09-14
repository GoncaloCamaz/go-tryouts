package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type AccountNotifier interface {
	// always a good practice to return an error, it might be useful in the future to handle failures
	NotifyAccountCreated(ctx context.Context, account Account) error
}

type SimpleAccountNotifier struct{}

func (n SimpleAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("new simple account created", "username", account.Username, "email", account.Email)
	return nil
}

type MahnaMahnaAccountNotifier struct{}

func (n MahnaMahnaAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("new Mahna Mahna account created", "username", account.Username, "email", account.Email)
	return nil
}

type Account struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

// What if, we wanted to notify the user that their account was created and we have a lot of providers to notify the user?
// func notifyAccountCreated(account Account) error {
// 	time.Sleep(time.Millisecond * 500)
// 	slog.Info("new account created", "username", account.Username, "email", account.Email)
// 	return nil
// }

func main() {
	mux := http.NewServeMux()

	accountHandler := &AccountHandler{
		AccountNotifier: MahnaMahnaAccountNotifier{},
	}

	mux.HandleFunc("POST /account", accountHandler.handleCreateAccount)

	http.ListenAndServe(":3000", mux)
}
