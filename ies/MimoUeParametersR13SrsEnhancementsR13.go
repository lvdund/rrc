package ies

import "rrc/utils"

// MIMO-UE-Parameters-r13-srs-Enhancements-r13 ::= ENUMERATED
type MimoUeParametersR13SrsEnhancementsR13 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParametersR13SrsEnhancementsR13EnumeratedNothing = iota
	MimoUeParametersR13SrsEnhancementsR13EnumeratedSupported
)
