package ies

import "rrc/utils"

// SL-RS-Type-r16 ::= ENUMERATED
type SlRsTypeR16 struct {
	Value utils.ENUMERATED
}

const (
	SlRsTypeR16EnumeratedNothing = iota
	SlRsTypeR16EnumeratedDmrs
	SlRsTypeR16EnumeratedSpare3
	SlRsTypeR16EnumeratedSpare2
	SlRsTypeR16EnumeratedSpare1
)
