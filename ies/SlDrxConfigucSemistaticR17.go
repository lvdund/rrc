package ies

import "rrc/utils"

// SL-DRX-ConfigUC-SemiStatic-r17 ::= SEQUENCE
type SlDrxConfigucSemistaticR17 struct {
	SlDrxOndurationtimerR17  SlDrxConfigucSemistaticR17SlDrxOndurationtimerR17
	SlDrxCyclestartoffsetR17 SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17
	SlDrxSlotoffsetR17       utils.INTEGER `lb:0,ub:31`
}
