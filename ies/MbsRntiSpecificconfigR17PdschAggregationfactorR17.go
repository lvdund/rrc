package ies

import "rrc/utils"

// MBS-RNTI-SpecificConfig-r17-pdsch-AggregationFactor-r17 ::= ENUMERATED
type MbsRntiSpecificconfigR17PdschAggregationfactorR17 struct {
	Value utils.ENUMERATED
}

const (
	MbsRntiSpecificconfigR17PdschAggregationfactorR17EnumeratedNothing = iota
	MbsRntiSpecificconfigR17PdschAggregationfactorR17EnumeratedN2
	MbsRntiSpecificconfigR17PdschAggregationfactorR17EnumeratedN4
	MbsRntiSpecificconfigR17PdschAggregationfactorR17EnumeratedN8
)
