package wallet

import (
	"encoding/json"
	"log"
	"net/http"

	d "github.com/bahodurnazarov/middleTask/db"
	"github.com/gorilla/mux"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	walletID := params["id"]

	db := d.Conn()
	query := "SELECT balance FROM wallets WHERE id = $1"
	row := db.QueryRow(query, walletID)

	var balance float64
	err := row.Scan(&balance)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	balanceResponse := map[string]interface{}{
		"balance": balance,
	}

	json.NewEncoder(w).Encode(balanceResponse)
}
