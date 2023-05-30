package wallet

import (
	"encoding/json"
	"net/http"

	d "github.com/bahodurnazarov/middleTask/db"
	lg "github.com/bahodurnazarov/middleTask/utils"
	"github.com/gorilla/mux"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	// Обработка GET-запроса
	params := mux.Vars(r)
	walletID := params["id"]

	db := d.Conn()
	query := "SELECT balance FROM wallets WHERE id = $1"
	row := db.QueryRow(query, walletID)

	var balance float64
	err := row.Scan(&balance)
	if err != nil {
		lg.Errl.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	balanceResponse := map[string]interface{}{
		"balance": balance,
	}

	json.NewEncoder(w).Encode(balanceResponse)
	lg.Server.Println(balanceResponse)
}
