package wallet

import (
	"encoding/json"
	"net/http"
	"time"

	d "github.com/bahodurnazarov/middleTask/db"
	lg "github.com/bahodurnazarov/middleTask/utils"
	"github.com/gorilla/mux"
)

type Transaction struct {
	ID              int       `json:"id"`
	WalletID        int       `json:"wallet_id"`
	Amount          float64   `json:"amount"`
	Description     string    `json:"description"`
	TransactionTime time.Time `json:"transaction_time"`
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	// Обработка POST-запроса
	params := mux.Vars(r)
	walletID := params["id"]

	userID := r.Header.Get("X-UserId")
	digest := r.Header.Get("X-Digest")

	// Verify authentication
	if AuthenticateRequest([]byte(userID), digest) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var transaction Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Retrieve wallet details
	wallet, err := GetWallet(walletID)
	if err != nil {
		lg.Errl.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Verify maximum balance based on identification status
	if transaction.Amount > 100000 && !wallet.Identified {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Update wallet balance
	newBalance := wallet.Balance + transaction.Amount
	if newBalance > 100000 && !wallet.Identified {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	db := d.Conn()
	_, err = db.Exec("UPDATE wallets SET balance = $1 WHERE id = $2", newBalance, wallet.ID)
	if err != nil {
		lg.Errl.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create transaction record
	_, err = db.Exec("INSERT INTO transactions (wallet_id, amount, description, transaction_time) VALUES ($1, $2, $3, $4)",
		wallet.ID, transaction.Amount, transaction.Description, time.Now())
	if err != nil {
		lg.Errl.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	lg.Server.Println(http.StatusOK)
}
