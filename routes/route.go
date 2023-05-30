package routes

import (
	"log"
	"net/http"

	"github.com/bahodurnazarov/middleTask/wallet"
	"github.com/gorilla/mux"
)

func Listening() {
	router := mux.NewRouter()
	router.HandleFunc("/wallets", wallet.CreateWallet).Methods("POST")
	router.HandleFunc("/wallets/{id}/exists", wallet.AccountExists).Methods("GET")
	router.HandleFunc("/wallets/{id}/deposit", wallet.Deposit).Methods("POST")
	router.HandleFunc("/wallets/{id}/summary", wallet.GetTransactionSummary).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}
