package ies

import "rrc/utils"

// UplinkPowerControlCommonSCell-r10 ::= SEQUENCE
type UplinkpowercontrolcommonscellR10 struct {
	P0NominalpuschR10 utils.INTEGER `lb:0,ub:24`
	AlphaR10          AlphaR12
}
