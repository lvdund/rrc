package ies

import "rrc/utils"

// NTN-Parameters-r17-ra-SDT-NTN-r17 ::= ENUMERATED
type NtnParametersR17RaSdtNtnR17 struct {
	Value utils.ENUMERATED
}

const (
	NtnParametersR17RaSdtNtnR17EnumeratedNothing = iota
	NtnParametersR17RaSdtNtnR17EnumeratedSupported
)
