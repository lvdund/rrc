package ies

import "rrc/utils"

// UplinkPowerControlCommonSCell-r10 ::= SEQUENCE
type UplinkpowercontrolcommonscellR10 struct {
	P0NominalpuschR10 utils.INTEGER
	AlphaR10          AlphaR12
}
