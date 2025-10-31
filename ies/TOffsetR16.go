package ies

import "rrc/utils"

// T-Offset-r16 ::= ENUMERATED
type TOffsetR16 struct {
	Value utils.ENUMERATED
}

const (
	TOffsetR16EnumeratedNothing = iota
	TOffsetR16EnumeratedMs0dot5
	TOffsetR16EnumeratedMs0dot75
	TOffsetR16EnumeratedMs1
	TOffsetR16EnumeratedMs1dot5
	TOffsetR16EnumeratedMs2
	TOffsetR16EnumeratedMs2dot5
	TOffsetR16EnumeratedMs3
	TOffsetR16EnumeratedSpare1
)
