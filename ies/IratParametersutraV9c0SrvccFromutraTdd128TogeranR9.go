package ies

import "rrc/utils"

// IRAT-ParametersUTRA-v9c0-srvcc-FromUTRA-TDD128-ToGERAN-r9 ::= ENUMERATED
type IratParametersutraV9c0SrvccFromutraTdd128TogeranR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersutraV9c0SrvccFromutraTdd128TogeranR9EnumeratedNothing = iota
	IratParametersutraV9c0SrvccFromutraTdd128TogeranR9EnumeratedSupported
)
