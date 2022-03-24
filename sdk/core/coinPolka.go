package sdk

type PolkaCoin struct {
	*baseCoin
}

func (c *PolkaCoin) AsPolka() *PolkaCoin {
	return c
}

func (c *PolkaCoin) QueryBalance() (string, error) {
	return "", nil
}
