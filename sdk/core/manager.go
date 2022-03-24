package sdk

type Manager struct {
}

var sharedManager *Manager

func GetSharedManager() *Manager {
	if sharedManager == nil {
		sharedManager = &Manager{}
	}
	return sharedManager
}

// 通过 CoinType 枚举获取 Coin 对象
// @return 如果币种不是内置的币，将返回空
func (m *Manager) CoinWithType(t CoinType) Coin {
	return internalCoinWithType(t)
}

// 通过 ChainType 枚举获取 Chain 对象
// @return 如果不是内置的链，将返回空
func (m *Manager) ChainWithType(t ChainType) Chain {
	return internalChainWithType(t)
}
