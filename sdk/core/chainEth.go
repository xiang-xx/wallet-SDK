package sdk

type EthChain struct {
	*baseChain
}

func NewEthChain(rpcUrl string) *EthChain {
	return &EthChain{
		baseChain: &baseChain{
			name:      ChainUserCustom,
			rpcUrl:    rpcUrl,
			framework: FrameworkEthereum,
		},
	}
}

func (c *EthChain) AsEth() *EthChain {
	return c
}

func (c *EthChain) EncodePubkey(pubkey string) (string, error) {
	return "", nil
}

func (c *EthChain) DecodeAddress(address string) (string, error) {
	return "", nil
}

// 获取合约币信息
func (c *EthChain) FetchContractCoin(contractAddress string) *EthCoin {
	return &EthCoin{&baseCoin{}, contractAddress}
}

// 主币的信息无法从链上获取
// 需要自己配置
func (c *EthChain) ConfigMainCoin(name string, symbol string, decimal int16) *EthCoin {
	return &EthCoin{&baseCoin{
		name:    name,
		symbol:  symbol,
		decimal: decimal,
	}, ""}
}
