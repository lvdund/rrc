package ies

import "rrc/utils"

// DC-Parameters-r12-drb-TypeSplit-r12 ::= ENUMERATED
type DcParametersR12DrbTypesplitR12 struct {
	Value utils.ENUMERATED
}

const (
	DcParametersR12DrbTypesplitR12EnumeratedNothing = iota
	DcParametersR12DrbTypesplitR12EnumeratedSupported
)
