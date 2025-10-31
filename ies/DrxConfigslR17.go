package ies

import "rrc/utils"

// DRX-ConfigSL-r17 ::= SEQUENCE
type DrxConfigslR17 struct {
	DrxHarqRttTimerslR17        utils.INTEGER `lb:0,ub:56`
	DrxRetransmissiontimerslR17 DrxConfigslR17DrxRetransmissiontimerslR17
}
