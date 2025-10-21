package ies

import "rrc/utils"

// UplinkPowerControlCommon-NB-r13 ::= SEQUENCE
type UplinkpowercontrolcommonNbR13 struct {
	P0NominalnpuschR13   utils.INTEGER
	AlphaR13             utils.ENUMERATED
	Deltapreamblemsg3R13 utils.INTEGER
}
