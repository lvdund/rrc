package ies

import "rrc/utils"

// LAA-Parameters-r13-tm10-LAA-r13 ::= ENUMERATED
type LaaParametersR13Tm10LaaR13 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersR13Tm10LaaR13EnumeratedNothing = iota
	LaaParametersR13Tm10LaaR13EnumeratedSupported
)
