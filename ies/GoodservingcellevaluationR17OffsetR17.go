package ies

import "rrc/utils"

// GoodServingCellEvaluation-r17-offset-r17 ::= ENUMERATED
type GoodservingcellevaluationR17OffsetR17 struct {
	Value utils.ENUMERATED
}

const (
	GoodservingcellevaluationR17OffsetR17EnumeratedNothing = iota
	GoodservingcellevaluationR17OffsetR17EnumeratedDb2
	GoodservingcellevaluationR17OffsetR17EnumeratedDb4
	GoodservingcellevaluationR17OffsetR17EnumeratedDb6
	GoodservingcellevaluationR17OffsetR17EnumeratedDb8
)
