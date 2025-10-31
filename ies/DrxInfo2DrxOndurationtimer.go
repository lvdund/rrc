package ies

import "rrc/utils"

// DRX-Info2-drx-onDurationTimer ::= CHOICE
const (
	DrxInfo2DrxOndurationtimerChoiceNothing = iota
	DrxInfo2DrxOndurationtimerChoiceSubmilliseconds
	DrxInfo2DrxOndurationtimerChoiceMilliseconds
)

type DrxInfo2DrxOndurationtimer struct {
	Choice          uint64
	Submilliseconds *utils.INTEGER `lb:1,ub:31`
	Milliseconds    *utils.ENUMERATED
}
