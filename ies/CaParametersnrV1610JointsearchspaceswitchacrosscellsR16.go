package ies

import "rrc/utils"

// CA-ParametersNR-v1610-jointSearchSpaceSwitchAcrossCells-r16 ::= ENUMERATED
type CaParametersnrV1610JointsearchspaceswitchacrosscellsR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610JointsearchspaceswitchacrosscellsR16EnumeratedNothing = iota
	CaParametersnrV1610JointsearchspaceswitchacrosscellsR16EnumeratedSupported
)
