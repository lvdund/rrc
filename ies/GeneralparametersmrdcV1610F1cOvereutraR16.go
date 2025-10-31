package ies

import "rrc/utils"

// GeneralParametersMRDC-v1610-f1c-OverEUTRA-r16 ::= ENUMERATED
type GeneralparametersmrdcV1610F1cOvereutraR16 struct {
	Value utils.ENUMERATED
}

const (
	GeneralparametersmrdcV1610F1cOvereutraR16EnumeratedNothing = iota
	GeneralparametersmrdcV1610F1cOvereutraR16EnumeratedSupported
)
