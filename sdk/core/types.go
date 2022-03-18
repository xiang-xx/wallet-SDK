package sdk

type SDKEnumInt = int
type SDKEnumString = string

type Account interface {
	// 公钥，以 0x 开头的
	Pubkey() string
	Address() string
}

type Tradable interface {
	payCoin() *Coin
	feeCoin() *Coin
}
