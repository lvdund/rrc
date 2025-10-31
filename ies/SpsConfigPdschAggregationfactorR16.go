package ies

import "rrc/utils"

// SPS-Config-pdsch-AggregationFactor-r16 ::= ENUMERATED
type SpsConfigPdschAggregationfactorR16 struct {
	Value utils.ENUMERATED
}

const (
	SpsConfigPdschAggregationfactorR16EnumeratedNothing = iota
	SpsConfigPdschAggregationfactorR16EnumeratedN1
	SpsConfigPdschAggregationfactorR16EnumeratedN2
	SpsConfigPdschAggregationfactorR16EnumeratedN4
	SpsConfigPdschAggregationfactorR16EnumeratedN8
)
