package ies

import "rrc/utils"

// LAA-Parameters-v1530-laa-PUSCH-Mode1-r15 ::= ENUMERATED
type LaaParametersV1530LaaPuschMode1R15 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersV1530LaaPuschMode1R15EnumeratedNothing = iota
	LaaParametersV1530LaaPuschMode1R15EnumeratedSupported
)
