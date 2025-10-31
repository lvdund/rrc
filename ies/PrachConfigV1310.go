package ies

import "rrc/utils"

// PRACH-Config-v1310 ::= SEQUENCE
type PrachConfigV1310 struct {
	RsrpThresholdsprachinfolistR13 *RsrpThresholdsprachinfolistR13
	MpdcchStartsfCssRaR13          *PrachConfigV1310MpdcchStartsfCssRaR13
	PrachHoppingoffsetR13          *utils.INTEGER `lb:0,ub:94`
	PrachParameterslistceR13       *PrachParameterslistceR13
	InitialCeLevelR13              *utils.INTEGER `lb:0,ub:3`
}
