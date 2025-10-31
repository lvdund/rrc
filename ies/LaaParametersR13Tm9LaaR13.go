package ies

import "rrc/utils"

// LAA-Parameters-r13-tm9-LAA-r13 ::= ENUMERATED
type LaaParametersR13Tm9LaaR13 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersR13Tm9LaaR13EnumeratedNothing = iota
	LaaParametersR13Tm9LaaR13EnumeratedSupported
)
