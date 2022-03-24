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

	framework FrameworkType
}

func internalChainWithType(t ChainType) Chain {
	rpc := rpcOfInternalChain(t)
	switch t {
	case ChainMinix, ChainChainx, ChainSherpax:
		return &PolkaChain{&baseChain{
			name:      t,
			rpcUrl:    rpc,
			framework: FrameworkPolka,
		}}
	case ChainSignet, ChainBitcoin:
		return &BtcChain{&baseChain{
			name:      t,
			rpcUrl:    rpc,
			framework: FrameworkBitcoin,
		}}
	case ChainEthereum, ChainSherpax_eth, ChainBinance:
		return &EthChain{&baseChain{
			name:      t,
			rpcUrl:    rpc,
			framework: FrameworkEthereum,
		}}
	}
	return nil
}

func rpcOfInternalChain(t ChainType) string {
	switch t {
	case ChainMinix:
	case ChainChainx:
	case ChainSherpax:
	case ChainSignet:
	case ChainBitcoin:
		return "https://..."
	case ChainEthereum:
		return "https://mainnet.infura.io/v3/da3717f25f824cc1baa32d812386d93f"
	case ChainSherpax_eth:
		return "https://mainnet.sherpax.io/rpc"
	case ChainBinance:
		return "https://bsc-dataseed.binance.org"
	}
	return ""
}

func (c *baseChain) Name() ChainType {
	return c.name
}
func (c *baseChain) RpcUrl() string {
	return c.rpcUrl
}
func (c *baseChain) Framework() FrameworkType {
	return c.framework
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

func (c *baseChain) AsEth() *EthChain {
	return nil
}
func (c *baseChain) AsPolka() *PolkaChain {
	return nil
}
func (c *baseChain) AsBtc() *BtcChain {
	return nil
}

// 公钥转地址
// @param pubkey 公钥，以 0x 开头
func (c *baseChain) EncodePubkey(pubkey string) (string, error) {
	return "", nil
}

// 地址转公钥
// @param address 地址
// @return 公钥，以 0x 开头
func (c *baseChain) DecodeAddress(address string) (string, error) {
	return "", nil
}
