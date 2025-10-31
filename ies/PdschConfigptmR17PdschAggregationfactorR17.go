package ies

import "rrc/utils"

// PDSCH-ConfigPTM-r17-pdsch-AggregationFactor-r17 ::= ENUMERATED
type PdschConfigptmR17PdschAggregationfactorR17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigptmR17PdschAggregationfactorR17EnumeratedNothing = iota
	PdschConfigptmR17PdschAggregationfactorR17EnumeratedN2
	PdschConfigptmR17PdschAggregationfactorR17EnumeratedN4
	PdschConfigptmR17PdschAggregationfactorR17EnumeratedN8
)
