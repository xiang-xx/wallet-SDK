package sdk

type ChainType = SDKEnumString

const (
	ChainMinix   ChainType = "minix" // Polka
	ChainChainx  ChainType = "chainx"
	ChainSherpax ChainType = "sherpax"

	ChainEthereum    ChainType = "ethereum" // Ethereum
	ChainSherpax_eth ChainType = "sherpax_eth"
	ChainBinance     ChainType = "binance"

	ChainBitcoin ChainType = "bitcoin" // Bitcoin
	ChainSignet  ChainType = "signet"

	ChainUserCustom ChainType = "CustomChain" // 用户自定义的链
)

type baseChain struct {
	name   ChainType
	rpcUrl string
}

func newChainWithType(t ChainType) *baseChain {
	rpc := ""
	switch t {
	case ChainMinix:
	case ChainChainx:
	case ChainSherpax:
	case ChainSignet:
	case ChainBitcoin:
		rpc = ""

	case ChainEthereum:
		rpc = "https://mainnet.infura.io/v3/da3717f25f824cc1baa32d812386d93f"
	case ChainSherpax_eth:
		rpc = "https://mainnet.sherpax.io/rpc"
	case ChainBinance:
		rpc = "https://bsc-dataseed.binance.org"

	default:
		return nil
	}

	return &baseChain{
		name:   t,
		rpcUrl: rpc,
	}
}

func (c *baseChain) Name() ChainType {
	return c.name
}
func (c *baseChain) RpcUrl() string {
	return c.rpcUrl
}
func (c *baseChain) Framework() FrameworkType {
	switch c.name {
	case ChainMinix, ChainChainx, ChainSherpax:
		return FrameworkPolka
	case ChainEthereum, ChainSherpax_eth, ChainBinance:
		return FrameworkEthereum
	case ChainSignet, ChainBitcoin:
		return FrameworkBitcoin
	default:
		return FrameworkUnknown
	}
}

func (c *baseChain) MainCoin() Coin {
	var cointype CoinType
	switch c.name {
	case ChainMinix:
		cointype = CoinMINI
	case ChainChainx:
		cointype = CoinPCX
	case ChainSherpax:
		cointype = CoinKSX

	case ChainEthereum:
		cointype = CoinETH
	case ChainSherpax_eth:
		cointype = CoinWKSX
	case ChainBinance:
		cointype = CoinBNB

	case ChainSignet:
		cointype = CoinSBTC
	case ChainBitcoin:
		cointype = CoinBTC
	default:
		return nil
	}
	return GetSharedManager().CoinWithType(cointype)
}

func (c *baseChain) CreateAccountWithMnemonic(mnemonic string) Wallet {
	// TODO:
	return &PersonalWallet{
		chain: c,
	}
}
func (c *baseChain) CreateAccountWithKeystore(keystore string, password string) Wallet {
	// TODO:
	return &PersonalWallet{
		chain: c,
	}
}
