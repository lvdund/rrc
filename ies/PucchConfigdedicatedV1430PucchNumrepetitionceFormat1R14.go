package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v1430-pucch-NumRepetitionCE-format1-r14 ::= ENUMERATED
type PucchConfigdedicatedV1430PucchNumrepetitionceFormat1R14 struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigdedicatedV1430PucchNumrepetitionceFormat1R14EnumeratedNothing = iota
	PucchConfigdedicatedV1430PucchNumrepetitionceFormat1R14EnumeratedR64
	PucchConfigdedicatedV1430PucchNumrepetitionceFormat1R14EnumeratedR128
)
