package ies

import "rrc/utils"

// LAA-Parameters-v1530-aul-r15 ::= ENUMERATED
type LaaParametersV1530AulR15 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersV1530AulR15EnumeratedNothing = iota
	LaaParametersV1530AulR15EnumeratedSupported
)
