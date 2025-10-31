package ies

import "rrc/utils"

// EUTRA-5GC-Parameters-v1610-ce-InactiveState-r16 ::= ENUMERATED
type Eutra5gcParametersV1610CeInactivestateR16 struct {
	Value utils.ENUMERATED
}

const (
	Eutra5gcParametersV1610CeInactivestateR16EnumeratedNothing = iota
	Eutra5gcParametersV1610CeInactivestateR16EnumeratedSupported
)
