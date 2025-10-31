package ies

import "rrc/utils"

// SL-DRX-ConfigUC-SemiStatic-r17-sl-drx-onDurationTimer-r17 ::= CHOICE
const (
	SlDrxConfigucSemistaticR17SlDrxOndurationtimerR17ChoiceNothing = iota
	SlDrxConfigucSemistaticR17SlDrxOndurationtimerR17ChoiceSubmilliseconds
	SlDrxConfigucSemistaticR17SlDrxOndurationtimerR17ChoiceMilliseconds
)

type SlDrxConfigucSemistaticR17SlDrxOndurationtimerR17 struct {
	Choice          uint64
	Submilliseconds *utils.INTEGER `lb:1,ub:31`
	Milliseconds    *utils.ENUMERATED
}
