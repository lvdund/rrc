package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-r13-csi-RS-EnhancementsTDD-r13 ::= ENUMERATED
type MimoUeParameterspertmR13CsiRsEnhancementstddR13 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmR13CsiRsEnhancementstddR13EnumeratedNothing = iota
	MimoUeParameterspertmR13CsiRsEnhancementstddR13EnumeratedSupported
)
