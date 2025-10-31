package ies

import "rrc/utils"

// DRX-ConfigPTM-r17-drx-onDurationTimerPTM-r17 ::= CHOICE
const (
	DrxConfigptmR17DrxOndurationtimerptmR17ChoiceNothing = iota
	DrxConfigptmR17DrxOndurationtimerptmR17ChoiceSubmilliseconds
	DrxConfigptmR17DrxOndurationtimerptmR17ChoiceMilliseconds
)

type DrxConfigptmR17DrxOndurationtimerptmR17 struct {
	Choice          uint64
	Submilliseconds *utils.INTEGER `lb:1,ub:31`
	Milliseconds    *utils.ENUMERATED
}
