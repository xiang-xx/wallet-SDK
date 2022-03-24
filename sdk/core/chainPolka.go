package sdk

type PolkaChain struct {
	*baseChain
}

func (c *PolkaChain) AsPolka() *PolkaChain {
	return c
}

func (c *PolkaChain) EncodePubkey(pubkey string) (string, error) {
	return "", nil
}

func (c *PolkaChain) DecodeAddress(address string) (string, error) {
	return "", nil
}
