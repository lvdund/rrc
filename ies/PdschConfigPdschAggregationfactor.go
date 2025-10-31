package ies

import "rrc/utils"

// PDSCH-Config-pdsch-AggregationFactor ::= ENUMERATED
type PdschConfigPdschAggregationfactor struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPdschAggregationfactorEnumeratedNothing = iota
	PdschConfigPdschAggregationfactorEnumeratedN2
	PdschConfigPdschAggregationfactorEnumeratedN4
	PdschConfigPdschAggregationfactorEnumeratedN8
)
