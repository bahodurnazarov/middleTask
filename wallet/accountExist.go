package wallet

import (
	"encoding/json"
	"log"
	"net/http"

	d "github.com/bahodurnazarov/middleTask/db"
	"github.com/gorilla/mux"
)

func AccountExists(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	walletID := params["id"]

	db := d.Conn()
	query := "SELECT COUNT(*) FROM wallets WHERE id = $1"
	row := db.QueryRow(query, walletID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	existsResponse := map[string]interface{}{
		"exists": count > 0,
	}

	json.NewEncoder(w).Encode(existsResponse)
}
