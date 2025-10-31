package ies

import "rrc/utils"

// LAA-Parameters-r13-endingDwPTS-r13 ::= ENUMERATED
type LaaParametersR13EndingdwptsR13 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersR13EndingdwptsR13EnumeratedNothing = iota
	LaaParametersR13EndingdwptsR13EnumeratedSupported
)
