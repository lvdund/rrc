package ies

import "rrc/utils"

// CA-ParametersNR-v1540 ::= SEQUENCE
type CaParametersnrV1540 struct {
	SimultaneoussrsAssoccsiRsAllcc         *utils.INTEGER `lb:0,ub:32`
	CsiRsImReceptionforfeedbackperbandcomb *CaParametersnrV1540CsiRsImReceptionforfeedbackperbandcomb
	SimultaneouscsiReportsallcc            *utils.INTEGER `lb:0,ub:32`
	DualpaArchitecture                     *CaParametersnrV1540DualpaArchitecture
}
