package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16-ce-MultiTB-EarlyTermination-r16 ::= ENUMERATED
type CeMultitbParametersR16CeMultitbEarlyterminationR16 struct {
	Value utils.ENUMERATED
}

const (
	CeMultitbParametersR16CeMultitbEarlyterminationR16EnumeratedNothing = iota
	CeMultitbParametersR16CeMultitbEarlyterminationR16EnumeratedSupported
)
