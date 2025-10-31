package ies

import "rrc/utils"

// PRACH-ConfigSIB-v1310 ::= SEQUENCE
type PrachConfigsibV1310 struct {
	RsrpThresholdsprachinfolistR13 RsrpThresholdsprachinfolistR13
	MpdcchStartsfCssRaR13          *PrachConfigsibV1310MpdcchStartsfCssRaR13
	PrachHoppingoffsetR13          *utils.INTEGER `lb:0,ub:94`
	PrachParameterslistceR13       PrachParameterslistceR13
}
