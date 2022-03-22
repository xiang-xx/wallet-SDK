package sdk

type PersonalWallet struct {
	chain Chain
}

// Wallet interface
func (p *PersonalWallet) Pubkey() string {
	// TODO
	return ""
}

// Wallet interface
func (p *PersonalWallet) Address() string {
	// TODO
	return ""
}

func (p *PersonalWallet) Chain() Chain {
	return p.chain
}

func (p *PersonalWallet) TransferAssetTo(recipient *Recipient, asset Asset) *Transaction {
	return &Transaction{
		Wallet:    p,
		Asset:     asset,
		Recipient: recipient,
	}
}
func (p *PersonalWallet) TransferAssetToMulti(multi *MultiRecipient, asset Asset) *Transaction {
	return &Transaction{
		Wallet:         p,
		Asset:          asset,
		MultiRecipient: multi,
	}
}
