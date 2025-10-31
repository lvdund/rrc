package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-ce-MultiTB-64QAM-r16 ::= ENUMERATED
type CeMultitbParametersR16CeMultitb64qamR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16CeMultitb64qamR16EnumeratedNothing = iota
	CeMultitbParametersR16CeMultitb64qamR16EnumeratedSupported
)
