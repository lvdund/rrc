package ies

import "rrc/utils"

// ValidityAreaList-r16 ::= SEQUENCE OF ValidityArea-r16
// SIZE (1..maxFreqIdle-r15)
type ValidityarealistR16 struct {
	Value utils.Sequence[ValidityareaR16]
}
