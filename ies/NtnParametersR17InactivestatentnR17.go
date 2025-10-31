package ies

import "rrc/utils"

// NTN-Parameters-r17-inactiveStateNTN-r17 ::= ENUMERATED
type NtnParametersR17InactivestatentnR17 struct {
	Value utils.ENUMERATED
}

const (
	NtnParametersR17InactivestatentnR17EnumeratedNothing = iota
	NtnParametersR17InactivestatentnR17EnumeratedSupported
)
