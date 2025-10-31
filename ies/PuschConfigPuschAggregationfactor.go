package ies

import "rrc/utils"

// PUSCH-Config-pusch-AggregationFactor ::= ENUMERATED
type PuschConfigPuschAggregationfactor struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigPuschAggregationfactorEnumeratedNothing = iota
	PuschConfigPuschAggregationfactorEnumeratedN2
	PuschConfigPuschAggregationfactorEnumeratedN4
	PuschConfigPuschAggregationfactorEnumeratedN8
)
