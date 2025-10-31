package ies

import "rrc/utils"

// CA-ParametersNR-v1610-scellDormancyOutsideActiveTime-r16 ::= ENUMERATED
type CaParametersnrV1610ScelldormancyoutsideactivetimeR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610ScelldormancyoutsideactivetimeR16EnumeratedNothing = iota
	CaParametersnrV1610ScelldormancyoutsideactivetimeR16EnumeratedSupported
)
