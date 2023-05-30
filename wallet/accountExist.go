package wallet

import (
	"encoding/json"
	"net/http"

	d "github.com/bahodurnazarov/middleTask/db"
	lg "github.com/bahodurnazarov/middleTask/utils"
	"github.com/gorilla/mux"
)

func AccountExists(w http.ResponseWriter, r *http.Request) {
	// Обработка GET-запроса
	params := mux.Vars(r)
	walletID := params["id"]
	lg.Server.Println(walletID)
	db := d.Conn()
	query := "SELECT COUNT(*) FROM wallets WHERE id = $1"
	row := db.QueryRow(query, walletID)

	var count int
	err := row.Scan(&count)
	if err != nil {
		lg.Errl.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	existsResponse := map[string]interface{}{
		"exists": count > 0,
	}

	json.NewEncoder(w).Encode(existsResponse)
	lg.Server.Println(existsResponse)
}
