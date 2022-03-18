package sdk

type Assert struct {
}

// Tradable interface
func (a *Assert) payCoin() *Coin {
	return NewCoinWithType(CoinPCX)
}

// Tradable interface
func (a *Assert) feeCoin() *Coin {
	return a.payCoin()
}
