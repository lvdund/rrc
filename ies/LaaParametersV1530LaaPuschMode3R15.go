package ies

import "rrc/utils"

// LAA-Parameters-v1530-laa-PUSCH-Mode3-r15 ::= ENUMERATED
type LaaParametersV1530LaaPuschMode3R15 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersV1530LaaPuschMode3R15EnumeratedNothing = iota
	LaaParametersV1530LaaPuschMode3R15EnumeratedSupported
)
