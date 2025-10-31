package ies

import "rrc/utils"

// MAC-Parameters-v1530-earlyData-UP-r15 ::= ENUMERATED
type MacParametersV1530EarlydataUpR15 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1530EarlydataUpR15EnumeratedNothing = iota
	MacParametersV1530EarlydataUpR15EnumeratedSupported
)
