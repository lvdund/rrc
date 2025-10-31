package ies

import "rrc/utils"

// DRX-Config ::= SEQUENCE
type DrxConfig struct {
	DrxOndurationtimer       DrxConfigDrxOndurationtimer
	DrxInactivitytimer       DrxConfigDrxInactivitytimer
	DrxHarqRttTimerdl        utils.INTEGER `lb:0,ub:56`
	DrxHarqRttTimerul        utils.INTEGER `lb:0,ub:56`
	DrxRetransmissiontimerdl DrxConfigDrxRetransmissiontimerdl
	DrxRetransmissiontimerul DrxConfigDrxRetransmissiontimerul
	DrxLongcyclestartoffset  DrxConfigDrxLongcyclestartoffset
	Shortdrx                 *DrxConfigShortdrx
	DrxSlotoffset            utils.INTEGER `lb:0,ub:31`
}
