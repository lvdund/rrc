package ies

import "rrc/utils"

// SL-Failure-r16-sl-Failure-r16 ::= ENUMERATED
type SlFailureR16SlFailureR16 struct {
	Value utils.ENUMERATED
}

const (
	SlFailureR16SlFailureR16EnumeratedNothing = iota
	SlFailureR16SlFailureR16EnumeratedRlf
	SlFailureR16SlFailureR16EnumeratedConfigfailure
	SlFailureR16SlFailureR16EnumeratedDrxreject_V1710
	SlFailureR16SlFailureR16EnumeratedSpare5
	SlFailureR16SlFailureR16EnumeratedSpare4
	SlFailureR16SlFailureR16EnumeratedSpare3
	SlFailureR16SlFailureR16EnumeratedSpare2
	SlFailureR16SlFailureR16EnumeratedSpare1
)
