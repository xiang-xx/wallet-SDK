package sdk

type BtcChain struct {
	*baseChain
}

func (c *BtcChain) AsBtc() *BtcChain {
	return c
}

func (c *BtcChain) EncodePubkey(pubkey string) (string, error) {
	return "", nil
}

func (c *BtcChain) DecodeAddress(address string) (string, error) {
	return "", nil
}
