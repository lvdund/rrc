package ies

import "rrc/utils"

// RSRQ-Type-r12 ::= SEQUENCE
type RsrqTypeR12 struct {
	AllsymbolsR12 utils.BOOLEAN
	WidebandR12   utils.BOOLEAN
}
