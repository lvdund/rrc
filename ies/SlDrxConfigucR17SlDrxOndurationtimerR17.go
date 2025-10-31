package ies

import "rrc/utils"

// SL-DRX-ConfigUC-r17-sl-drx-onDurationTimer-r17 ::= CHOICE
const (
	SlDrxConfigucR17SlDrxOndurationtimerR17ChoiceNothing = iota
	SlDrxConfigucR17SlDrxOndurationtimerR17ChoiceSubmilliseconds
	SlDrxConfigucR17SlDrxOndurationtimerR17ChoiceMilliseconds
)

type SlDrxConfigucR17SlDrxOndurationtimerR17 struct {
	Choice          uint64
	Submilliseconds *utils.INTEGER `lb:1,ub:31`
	Milliseconds    *utils.ENUMERATED
}
