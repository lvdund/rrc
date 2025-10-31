package ies

import "rrc/utils"

// SL-DRX-ConfigUC-r17 ::= SEQUENCE
type SlDrxConfigucR17 struct {
	SlDrxOndurationtimerR17     SlDrxConfigucR17SlDrxOndurationtimerR17
	SlDrxInactivitytimerR17     SlDrxConfigucR17SlDrxInactivitytimerR17
	SlDrxHarqRttTimer1R17       *SlDrxConfigucR17SlDrxHarqRttTimer1R17
	SlDrxHarqRttTimer2R17       *SlDrxConfigucR17SlDrxHarqRttTimer2R17
	SlDrxRetransmissiontimerR17 SlDrxConfigucR17SlDrxRetransmissiontimerR17
	SlDrxCyclestartoffsetR17    SlDrxConfigucR17SlDrxCyclestartoffsetR17
	SlDrxSlotoffset             utils.INTEGER `lb:0,ub:31`
}
