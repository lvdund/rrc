package ies

import "rrc/utils"

// DRX-ConfigPTM-r17 ::= SEQUENCE
type DrxConfigptmR17 struct {
	DrxOndurationtimerptmR17       DrxConfigptmR17DrxOndurationtimerptmR17
	DrxInactivitytimerptmR17       DrxConfigptmR17DrxInactivitytimerptmR17
	DrxHarqRttTimerdlPtmR17        *utils.INTEGER `lb:0,ub:56`
	DrxRetransmissiontimerdlPtmR17 *DrxConfigptmR17DrxRetransmissiontimerdlPtmR17
	DrxLongcyclestartoffsetptmR17  DrxConfigptmR17DrxLongcyclestartoffsetptmR17
	DrxSlotoffsetptmR17            utils.INTEGER `lb:0,ub:31`
}
