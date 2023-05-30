package wallet

import (
	"encoding/json"
	"net/http"
	"time"

	d "github.com/bahodurnazarov/middleTask/db"
	lg "github.com/bahodurnazarov/middleTask/utils"
	"github.com/gorilla/mux"
)

func GetTransactionSummary(w http.ResponseWriter, r *http.Request) {

	// Обработка GET-запроса
	params := mux.Vars(r)
	walletID := params["id"]

	currentYear, currentMonth, _ := time.Now().Date()

	db := d.Conn()
	query := "SELECT COUNT(*), COALESCE(SUM(amount), 0) FROM transactions WHERE wallet_id = $1 AND EXTRACT(YEAR FROM transaction_time) = $2 AND EXTRACT(MONTH FROM transaction_time) = $3"
	row := db.QueryRow(query, walletID, currentYear, int(currentMonth))

	var count int
	var sum float64
	err := row.Scan(&count, &sum)
	if err != nil {
		lg.Errl.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	summary := map[string]interface{}{
		"count": count,
		"sum":   sum,
	}

	json.NewEncoder(w).Encode(summary)
	lg.Server.Println(summary)
}
