package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-ce-MultiTB-Interleaving-r16 ::= ENUMERATED
type CeMultitbParametersR16CeMultitbInterleavingR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16CeMultitbInterleavingR16EnumeratedNothing = iota
	CeMultitbParametersR16CeMultitbInterleavingR16EnumeratedSupported
)
