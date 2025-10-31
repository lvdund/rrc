package ies

import "rrc/utils"

// LWIP-Parameters-v1430-lwip-Aggregation-DL-r14 ::= ENUMERATED
type LwipParametersV1430LwipAggregationDlR14 struct {
	Value utils.ENUMERATED
}

const (
	LwipParametersV1430LwipAggregationDlR14EnumeratedNothing = iota
	LwipParametersV1430LwipAggregationDlR14EnumeratedSupported
)
