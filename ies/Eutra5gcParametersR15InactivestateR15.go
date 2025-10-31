package ies

import "rrc/utils"

// EUTRA-5GC-Parameters-r15-inactiveState-r15 ::= ENUMERATED
type Eutra5gcParametersR15InactivestateR15 struct {
	Value utils.ENUMERATED
}

const (
	Eutra5gcParametersR15InactivestateR15EnumeratedNothing = iota
	Eutra5gcParametersR15InactivestateR15EnumeratedSupported
)
