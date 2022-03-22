package sdk

type Transaction struct {
	Wallet Wallet
	Asset  Asset
	// 转帐给个人
	Recipient *Recipient
	// 转账给多个对象
	MultiRecipient *MultiRecipient
}

func (t *Transaction) EstimateFee() (string, error) {
	return "0", nil
}

func (t *Transaction) SignRawTx(password string) (string, error) {
	return "", nil
}
