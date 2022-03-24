package sdk

type EthCoin struct {
	*baseCoin

	contractAddress string // 合约地址
}

func (c *EthCoin) AsEth() *EthCoin {
	return c
}

func (c *EthCoin) QueryBalance() (string, error) {
	return "", nil
}

// 合约币的合约地址
func (c *EthCoin) ContractAddress() string {
	return c.contractAddress
}

func contractOfInternalEthCoin(t CoinType) string {
	switch t {
	case CoinUSDT:
		return "0xdac17f958d2ee523a2206206994597c13d831ec7"

	case CoinWKSX_USB:
		return "0xE7e312dfC08e060cda1AF38C234AEAcc7A982143"
	case CoinWKSX_USDT:
		return "0x4B53739D798EF0BEa5607c254336b40a93c75b52"
	case CoinWKSX_BUSD:
		return "0x37088186089c7D6BcD556d9A15087DFaE3Ba0C32"
	case CoinWKSX_USDC:
		return "0x935CC842f220CF3A7D10DA1c99F01B1A6894F7C5"

	case CoinBSC_USDT:
		return "0x55d398326f99059fF775485246999027B3197955"
	case CoinBSC_BUSD:
		return "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"
	case CoinBSC_USDC:
		return "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d"
	}
	return ""
}
