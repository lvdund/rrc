package ies

// RLC-Parameters-v1530 ::= SEQUENCE
type RlcParametersV1530 struct {
	FlexibleumAmCombinationsR15 *RlcParametersV1530FlexibleumAmCombinationsR15
	RlcAmOooDeliveryR15         *RlcParametersV1530RlcAmOooDeliveryR15
	RlcUmOooDeliveryR15         *RlcParametersV1530RlcUmOooDeliveryR15
}
