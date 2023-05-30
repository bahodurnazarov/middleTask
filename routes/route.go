package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	lg "github.com/bahodurnazarov/middleTask/utils"
	"github.com/bahodurnazarov/middleTask/wallet"
)

func Listening() {
	router := mux.NewRouter()
	router.HandleFunc("/wallets", wallet.CreateWallet).Methods("POST")
	router.HandleFunc("/wallets/{id}/exists", wallet.AccountExists).Methods("GET")
	router.HandleFunc("/wallets/{id}/deposit", wallet.Deposit).Methods("POST")
	router.HandleFunc("/wallets/{id}/summary", wallet.GetTransactionSummary).Methods("GET")
	router.HandleFunc("/wallets/{id}/balance", wallet.GetBalance).Methods("GET")
	fmt.Println("Сервер запущен на http://localhost:8000")
	lg.Server.Fatal(http.ListenAndServe(":8000", router))
}
