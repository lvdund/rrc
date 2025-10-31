package ies

import "rrc/utils"

// FailureInfoDAPS-r16-failureType-r16 ::= ENUMERATED
type FailureinfodapsR16FailuretypeR16 struct {
	Value utils.ENUMERATED
}

const (
	FailureinfodapsR16FailuretypeR16EnumeratedNothing = iota
	FailureinfodapsR16FailuretypeR16EnumeratedDaps_Failure
	FailureinfodapsR16FailuretypeR16EnumeratedSpare3
	FailureinfodapsR16FailuretypeR16EnumeratedSpare2
	FailureinfodapsR16FailuretypeR16EnumeratedSpare1
)
