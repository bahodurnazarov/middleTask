package wallet

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	lg "github.com/bahodurnazarov/middleTask/utils"
	d "github.com/bahodurnazarov/middleTask/db"
)

type Wallet struct {
	ID           int
	UserID       string
	Balance      float64
	Identified   bool
	CreationTime time.Time
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	// Обработка POST-запроса// Чтение данных запроса
	var wallet Wallet
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Проверка наличия обязательных полей
	if wallet.UserID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Проверка максимального баланса
	if wallet.Balance > 100000 && !wallet.Identified {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	db := d.Conn()
	// Вставка записи в базу данных
	query := "INSERT INTO wallets (user_id, balance, identified, creation_time) VALUES ($1, $2, $3, $4) RETURNING id"
	err = db.QueryRow(query, wallet.UserID, wallet.Balance, wallet.Identified, time.Now()).Scan(&wallet.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Отправка ответа со статусом 201 Created и информацией о созданном кошельке
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wallet)
	lg.Server.Println(wallet)
}
