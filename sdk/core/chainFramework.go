package sdk

type FrameworkType = SDKEnumString

const (
	FrameworkPolka    FrameworkType = "polka"
	FrameworkBitcoin  FrameworkType = "bitcoin"
	FrameworkEthereum FrameworkType = "ethereum"

	FrameworkUnknown FrameworkType = "unknown"
)
