package ies

import "rrc/utils"

// UplinkPowerControlCommon-NB-r13 ::= SEQUENCE
type UplinkpowercontrolcommonNbR13 struct {
	P0NominalnpuschR13   utils.INTEGER `lb:0,ub:24`
	AlphaR13             UplinkpowercontrolcommonNbR13AlphaR13
	Deltapreamblemsg3R13 utils.INTEGER `lb:0,ub:6`
}
