package sdk

import "strings"

type CoinType = SDKEnumString

const (
	CoinMINI CoinType = "MINI" // polka
	CoinPCX  CoinType = "PCX"
	CoinXBTC CoinType = "XBTC"
	CoinKSX  CoinType = "KSX"

	CoinETH  CoinType = "ETH" // ETH
	CoinUSDT CoinType = "USDT"

	CoinSBTC CoinType = "sBTC" // sBTC
	CoinBTC  CoinType = "BTC"  // BTC

	CoinWKSX      CoinType = "WKSX" // WKSX
	CoinWKSX_USB  CoinType = "WKSX_USB"
	CoinWKSX_USDT CoinType = "WKSX_USDT"
	CoinWKSX_BUSD CoinType = "WKSX_BUSD"
	CoinWKSX_USDC CoinType = "WKSX_USDC"

	CoinBNB      CoinType = "BNB" // Binance
	CoinBSC_USDT CoinType = "BSC_USDT"
	CoinBSC_BUSD CoinType = "BSC_BUSD"
	CoinBSC_USDC CoinType = "BSC_USDC"

	CoinUserCustom CoinType = "CustomCoin" // 用户自定义的币
)

type baseCoin struct {
	name    CoinType // 币种唯一标识符
	symbol  string   // 币种符号，可展示给用户看
	decimal int16    // 精度

	contractAddress string // 合约地址
}

// 通过 CoinType 枚举，来创建币种对象
// @return 如果币种不是内置的币，将返回空
func newCoinWithType(t CoinType) *baseCoin {
	var name, symbol string

	upperName := strings.ToUpper(t)
	switch upperName {
	case CoinMINI,
		CoinPCX,
		CoinXBTC,
		CoinETH,
		CoinUSDT,
		CoinKSX,
		CoinWKSX,
		CoinBNB,
		CoinBTC:
		name = upperName
		symbol = upperName

	case "SBTC": // sBTC 首字母需要小写
		name = CoinSBTC
		symbol = CoinSBTC

	case CoinWKSX_USB,
		CoinWKSX_USDT,
		CoinWKSX_BUSD,
		CoinWKSX_USDC,
		CoinBSC_USDT,
		CoinBSC_BUSD,
		CoinBSC_USDC:
		name = upperName
		symbol = strings.Split(upperName, "_")[1]
	default:
		return nil
	}

	return &baseCoin{
		name:   name,
		symbol: symbol,
	}
}

func (c *baseCoin) PayCoin() Coin {
	return c
}

func (c *baseCoin) FeeCoin() Coin {
	return c.Chain().MainCoin()
}

// properties

// 币种唯一标识符
func (c *baseCoin) Name() CoinType {
	return c.name
}

// 币种符号，可展示给用户看
func (c *baseCoin) Symbol() string {
	return c.symbol
}

func (c *baseCoin) Decimal() int16 {
	return c.decimal
}

func (c *baseCoin) Chain() Chain {
	return nil
}

// 是否是链上的主币
func (c *baseCoin) IsMainCoin() bool {
	return c.Chain().MainCoin().Name() == c.Name()
}

// 合约币的合约地址
func (c *baseCoin) ContractAddress() string {
	if c.contractAddress != "" {
		return c.contractAddress
	}
	switch c.name {
	// case CoinXBTC: // ??
	case CoinUSDT:
		c.contractAddress = "0xdac17f958d2ee523a2206206994597c13d831ec7"
	case CoinWKSX_USB:
		c.contractAddress = "0xE7e312dfC08e060cda1AF38C234AEAcc7A982143"
	case CoinWKSX_USDT:
		c.contractAddress = "0x4B53739D798EF0BEa5607c254336b40a93c75b52"
	case CoinWKSX_BUSD:
		c.contractAddress = "0x37088186089c7D6BcD556d9A15087DFaE3Ba0C32"
	case CoinWKSX_USDC:
		c.contractAddress = "0x935CC842f220CF3A7D10DA1c99F01B1A6894F7C5"
	case CoinBSC_USDT:
		c.contractAddress = "0x55d398326f99059fF775485246999027B3197955"
	case CoinBSC_BUSD:
		c.contractAddress = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"
	case CoinBSC_USDC:
		c.contractAddress = "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d"
	}
	return c.contractAddress
}

// 修改合约地址，提供 setter 只是为了方便 debug
func (c *baseCoin) SetContractAddress(address string) {
	c.contractAddress = address
}
