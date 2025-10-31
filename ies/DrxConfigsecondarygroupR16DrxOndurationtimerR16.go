package ies

import "rrc/utils"

// DRX-ConfigSecondaryGroup-r16-drx-onDurationTimer-r16 ::= CHOICE
const (
	DrxConfigsecondarygroupR16DrxOndurationtimerR16ChoiceNothing = iota
	DrxConfigsecondarygroupR16DrxOndurationtimerR16ChoiceSubmilliseconds
	DrxConfigsecondarygroupR16DrxOndurationtimerR16ChoiceMilliseconds
)

type DrxConfigsecondarygroupR16DrxOndurationtimerR16 struct {
	Choice          uint64
	Submilliseconds *utils.INTEGER `lb:1,ub:31`
	Milliseconds    *utils.ENUMERATED
}
