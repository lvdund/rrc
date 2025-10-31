package ies

import "rrc/utils"

// FailureInformation-r16-IEs-failureType-r16 ::= ENUMERATED
type FailureinformationR16IesFailuretypeR16 struct {
	Value utils.ENUMERATED
}

const (
	FailureinformationR16IesFailuretypeR16EnumeratedNothing = iota
	FailureinformationR16IesFailuretypeR16EnumeratedDuplication
	FailureinformationR16IesFailuretypeR16EnumeratedDapsho_Failure
	FailureinformationR16IesFailuretypeR16EnumeratedSpare2
	FailureinformationR16IesFailuretypeR16EnumeratedSpare1
)
