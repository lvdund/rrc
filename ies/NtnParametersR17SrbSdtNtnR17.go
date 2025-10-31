package ies

import "rrc/utils"

// NTN-Parameters-r17-srb-SDT-NTN-r17 ::= ENUMERATED
type NtnParametersR17SrbSdtNtnR17 struct {
	Value utils.ENUMERATED
}

const (
	NtnParametersR17SrbSdtNtnR17EnumeratedNothing = iota
	NtnParametersR17SrbSdtNtnR17EnumeratedSupported
)
