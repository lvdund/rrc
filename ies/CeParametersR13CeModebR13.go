package ies

import "rrc/utils"

// CE-Parameters-r13-ce-ModeB-r13 ::= ENUMERATED
type CeParametersR13CeModebR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersR13CeModebR13EnumeratedNothing = iota
	CeParametersR13CeModebR13EnumeratedSupported
)
