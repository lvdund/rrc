package ies

import "rrc/utils"

// IRAT-ParametersUTRA-v9c0-srvcc-FromUTRA-TDD128-ToUTRA-TDD128-r9 ::= ENUMERATED
type IratParametersutraV9c0SrvccFromutraTdd128ToutraTdd128R9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersutraV9c0SrvccFromutraTdd128ToutraTdd128R9EnumeratedNothing = iota
	IratParametersutraV9c0SrvccFromutraTdd128ToutraTdd128R9EnumeratedSupported
)
