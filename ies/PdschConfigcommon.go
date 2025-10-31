package ies

import "rrc/utils"

// PDSCH-ConfigCommon ::= SEQUENCE
type PdschConfigcommon struct {
	Referencesignalpower utils.INTEGER `lb:0,ub:50`
	PB                   utils.INTEGER `lb:0,ub:3`
}
