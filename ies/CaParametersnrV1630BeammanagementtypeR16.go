package ies

import "rrc/utils"

// CA-ParametersNR-v1630-beamManagementType-r16 ::= ENUMERATED
type CaParametersnrV1630BeammanagementtypeR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1630BeammanagementtypeR16EnumeratedNothing = iota
	CaParametersnrV1630BeammanagementtypeR16EnumeratedIbm
	CaParametersnrV1630BeammanagementtypeR16EnumeratedDummy
)
