package ies

import "rrc/utils"

// PUSCH-ConfigCommon-pusch-ConfigBasic ::= SEQUENCE
type PuschConfigcommonPuschConfigbasic struct {
	NSb                utils.INTEGER `lb:0,ub:4`
	Hoppingmode        PuschConfigcommonPuschConfigbasicHoppingmode
	PuschHoppingoffset utils.INTEGER `lb:0,ub:98`
	Enable64qam        utils.BOOLEAN
}
