package sdk

type ChainType = SDKEnumString

const (
	ChainMinix       ChainType = "minix"
	ChainChainx      ChainType = "chainx"
	ChainEthereum    ChainType = "ethereum"
	ChainSignet      ChainType = "signet"
	ChainSherpax     ChainType = "sherpax"
	ChainSherpax_eth ChainType = "sherpax_eth"
	ChainBinance     ChainType = "binance"
	ChainBitcoin     ChainType = "bitcoin"

	// 用户自定义的链
	ChainUserCustom ChainType = "CustomChain"
)

type Chain struct {
	Name      ChainType
	RpcUrl    string
	Framework ChainFramework
}
