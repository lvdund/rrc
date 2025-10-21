package ies

import "rrc/utils"

// LWIP-Parameters-v1430 ::= SEQUENCE
type LwipParametersV1430 struct {
	LwipAggregationDlR14 *utils.ENUMERATED
	LwipAggregationUlR14 *utils.ENUMERATED
}
