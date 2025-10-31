package ies

import "rrc/utils"

// CE-Parameters-r13-ce-ModeA-r13 ::= ENUMERATED
type CeParametersR13CeModeaR13 struct {
	Value utils.ENUMERATED
}

const (
	CeParametersR13CeModeaR13EnumeratedNothing = iota
	CeParametersR13CeModeaR13EnumeratedSupported
)
