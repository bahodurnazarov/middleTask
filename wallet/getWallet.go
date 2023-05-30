package wallet

import d "github.com/bahodurnazarov/middleTask/db"

func GetWallet(walletID string) (*Wallet, error) {

	db := d.Conn()
	query := "SELECT id, user_id, balance, identified, creation_time FROM wallets WHERE id = $1"
	row := db.QueryRow(query, walletID)

	var wallet Wallet
	err := row.Scan(&wallet.ID, &wallet.UserID, &wallet.Balance, &wallet.Identified, &wallet.CreationTime)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
