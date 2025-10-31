package ies

import "rrc/utils"

// SystemInformationBlockType14-r11 ::= SEQUENCE
// Extensible
type Systeminformationblocktype14R11 struct {
	EabParamR11              *Systeminformationblocktype14R11EabParamR11
	Latenoncriticalextension *utils.OCTETSTRING
}
