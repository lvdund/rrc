package ies

import "rrc/utils"

// DRX-Config-NB-r13-setup ::= SEQUENCE
type DrxConfigNbR13Setup struct {
	OndurationtimerR13          DrxConfigNbR13SetupOndurationtimerR13
	DrxInactivitytimerR13       DrxConfigNbR13SetupDrxInactivitytimerR13
	DrxRetransmissiontimerR13   DrxConfigNbR13SetupDrxRetransmissiontimerR13
	DrxCycleR13                 DrxConfigNbR13SetupDrxCycleR13
	DrxStartoffsetR13           utils.INTEGER `lb:0,ub:255`
	DrxUlretransmissiontimerR13 DrxConfigNbR13SetupDrxUlretransmissiontimerR13
}
