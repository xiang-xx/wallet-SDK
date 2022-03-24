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
}

// 通过 CoinType 枚举，来创建币种对象
// @return 如果币种不是内置的币，将返回空
func internalCoinWithType(t CoinType) Coin {
	var name, symbol string

	upperName := strings.ToUpper(t)
	switch upperName {
	case CoinMINI,
		CoinPCX,
		CoinXBTC,
		CoinKSX:
		return &PolkaCoin{&baseCoin{
			name:   upperName,
			symbol: upperName,
		}}

	case CoinBTC:
		return &BtcCoin{&baseCoin{
			name:   upperName,
			symbol: upperName,
		}}
	case "SBTC": // sBTC 首字母是小写的
		return &BtcCoin{&baseCoin{
			name:   CoinSBTC,
			symbol: CoinSBTC,
		}}

	case CoinETH,
		CoinUSDT,
		CoinWKSX,
		CoinBNB:
		contractAddress := contractOfInternalEthCoin(t)
		return &EthCoin{&baseCoin{
			name:   upperName,
			symbol: upperName,
		}, contractAddress}
	case CoinWKSX_USB,
		CoinWKSX_USDT,
		CoinWKSX_BUSD,
		CoinWKSX_USDC,
		CoinBSC_USDT,
		CoinBSC_BUSD,
		CoinBSC_USDC:
		name = upperName
		symbol = strings.Split(upperName, "_")[1]
		contractAddress := contractOfInternalEthCoin(t)
		return &EthCoin{&baseCoin{
			name:   name,
			symbol: symbol,
		}, contractAddress}
	}

	return nil
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

func (c *baseCoin) AsEth() *EthCoin {
	return nil
}
func (c *baseCoin) AsPolka() *PolkaCoin {
	return nil
}
func (c *baseCoin) AsBtc() *BtcCoin {
	return nil
}
func (c *baseCoin) QueryBalance() (string, error) {
	return "", nil
}
