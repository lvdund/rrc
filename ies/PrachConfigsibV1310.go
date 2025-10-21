package ies

import "rrc/utils"

// PRACH-ConfigSIB-v1310 ::= SEQUENCE
type PrachConfigsibV1310 struct {
	RsrpThresholdsprachinfolistR13 RsrpThresholdsprachinfolistR13
	MpdcchStartsfCssRaR13          *interface{}
	PrachHoppingoffsetR13          *utils.INTEGER
	PrachParameterslistceR13       PrachParameterslistceR13
}
