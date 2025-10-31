package ies

import "rrc/utils"

// CA-ParametersNR-v1610-scellDormancyWithinActiveTime-r16 ::= ENUMERATED
type CaParametersnrV1610ScelldormancywithinactivetimeR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610ScelldormancywithinactivetimeR16EnumeratedNothing = iota
	CaParametersnrV1610ScelldormancywithinactivetimeR16EnumeratedSupported
)
