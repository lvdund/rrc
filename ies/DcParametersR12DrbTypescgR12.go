package ies

import "rrc/utils"

// DC-Parameters-r12-drb-TypeSCG-r12 ::= ENUMERATED
type DcParametersR12DrbTypescgR12 struct {
	Value utils.ENUMERATED
}

const (
	DcParametersR12DrbTypescgR12EnumeratedNothing = iota
	DcParametersR12DrbTypescgR12EnumeratedSupported
)
