package ies

// ValidityAreaList-r16 ::= SEQUENCE OF ValidityArea-r16
// SIZE (1..maxFreqIdle-r15)
type ValidityarealistR16 struct {
	Value []ValidityareaR16 `lb:1,ub:maxFreqIdleR15`
}
