package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-ce-MultiTB-SubPRB-r16 ::= ENUMERATED
type CeMultitbParametersR16CeMultitbSubprbR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16CeMultitbSubprbR16EnumeratedNothing = iota
	CeMultitbParametersR16CeMultitbSubprbR16EnumeratedSupported
)
