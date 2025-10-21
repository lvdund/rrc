package ies

import "rrc/utils"

// RLC-Parameters-v1530 ::= SEQUENCE
type RlcParametersV1530 struct {
	FlexibleumAmCombinationsR15 *utils.ENUMERATED
	RlcAmOooDeliveryR15         *utils.ENUMERATED
	RlcUmOooDeliveryR15         *utils.ENUMERATED
}
