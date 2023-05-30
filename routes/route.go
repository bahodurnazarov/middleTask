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

	log.Fatal(http.ListenAndServe(":8000", router))
}
