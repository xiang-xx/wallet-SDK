package sdk

type SDKEnumInt = int
type SDKEnumString = string

type Wallet interface {
	// 公钥，以 0x 开头的
	Pubkey() string
	Address() string
	// 钱包所属的链
	Chain() Chain

	// 给单人转账
	TransferAssetTo(recipient *Recipient, asset Asset) *Transaction
	// 批量给多人转账
	TransferAssetToMulti(multi *MultiRecipient, asset Asset) *Transaction
}

type Asset interface {
	// 交易资产需要消耗的币
	PayCoin() Coin
	// 交易资产需要消耗的手续费币
	FeeCoin() Coin
	// 资产所属的链
	Chain() Chain
}

type Chain interface {
	Name() ChainType
	RpcUrl() string
	// 链的生态
	Framework() FrameworkType

	// 该链的主币
	MainCoin() Coin

	// 使用助记词创建钱包
	CreateAccountWithMnemonic(mnemonic string) Wallet
	// 使用 keystore 和密码创建钱包
	CreateAccountWithKeystore(keystore string, password string) Wallet
}

// 继承于 Asset, 可以当 Asset 对象来使用
type Coin interface {
	Asset
	// 币种的唯一标识符
	Name() CoinType
	// 币种的符号
	Symbol() string
	// 币种的精度
	Decimal() int16

	// 是否是主币
	IsMainCoin() bool
}
