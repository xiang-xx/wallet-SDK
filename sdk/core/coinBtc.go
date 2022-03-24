package sdk

type BtcCoin struct {
	*baseCoin
}

func (c *BtcCoin) AsBtc() *BtcCoin {
	return c
}

func (c *BtcCoin) QueryBalance() (string, error) {
	return "", nil
}
