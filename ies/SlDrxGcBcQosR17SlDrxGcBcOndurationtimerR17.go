package ies

import "rrc/utils"

// SL-DRX-GC-BC-QoS-r17-sl-DRX-GC-BC-OnDurationTimer-r17 ::= CHOICE
const (
	SlDrxGcBcQosR17SlDrxGcBcOndurationtimerR17ChoiceNothing = iota
	SlDrxGcBcQosR17SlDrxGcBcOndurationtimerR17ChoiceSubmilliseconds
	SlDrxGcBcQosR17SlDrxGcBcOndurationtimerR17ChoiceMilliseconds
)

type SlDrxGcBcQosR17SlDrxGcBcOndurationtimerR17 struct {
	Choice          uint64
	Submilliseconds *utils.INTEGER `lb:1,ub:31`
	Milliseconds    *utils.ENUMERATED
}
