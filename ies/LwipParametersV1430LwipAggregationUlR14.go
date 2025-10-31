package ies

import "rrc/utils"

// LWIP-Parameters-v1430-lwip-Aggregation-UL-r14 ::= ENUMERATED
type LwipParametersV1430LwipAggregationUlR14 struct {
	Value utils.ENUMERATED
}

const (
	LwipParametersV1430LwipAggregationUlR14EnumeratedNothing = iota
	LwipParametersV1430LwipAggregationUlR14EnumeratedSupported
)
