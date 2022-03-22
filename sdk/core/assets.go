package sdk

type NFT struct {
}

func (a *NFT) PayCoin() Coin {
	return a.Chain().MainCoin()
}

func (a *NFT) FeeCoin() Coin {
	return a.Chain().MainCoin()
}

func (a *NFT) Chain() Chain {
	return &baseChain{}
}
