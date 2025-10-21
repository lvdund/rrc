package ies

import "rrc/utils"

// RACH-ConfigCommonSCell-r11 ::= SEQUENCE
// Extensible
type RachConfigcommonscellR11 struct {
	PowerrampingparametersR11 Powerrampingparameters
	RaSupervisioninfoR11      interface{}
}
