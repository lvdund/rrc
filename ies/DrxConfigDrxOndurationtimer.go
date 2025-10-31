package ies

import "rrc/utils"

// DRX-Config-drx-onDurationTimer ::= CHOICE
const (
	DrxConfigDrxOndurationtimerChoiceNothing = iota
	DrxConfigDrxOndurationtimerChoiceSubmilliseconds
	DrxConfigDrxOndurationtimerChoiceMilliseconds
)

type DrxConfigDrxOndurationtimer struct {
	Choice          uint64
	Submilliseconds *utils.INTEGER `lb:1,ub:31`
	Milliseconds    *utils.ENUMERATED
}
