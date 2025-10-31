package ies

import "rrc/utils"

// MIMO-UE-Parameters-r13-srs-EnhancementsTDD-r13 ::= ENUMERATED
type MimoUeParametersR13SrsEnhancementstddR13 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParametersR13SrsEnhancementstddR13EnumeratedNothing = iota
	MimoUeParametersR13SrsEnhancementstddR13EnumeratedSupported
)
