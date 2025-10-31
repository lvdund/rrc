package ies

import "rrc/utils"

// NPDSCH-ConfigCommon-NB-r13 ::= SEQUENCE
type NpdschConfigcommonNbR13 struct {
	NrsPowerR13 utils.INTEGER `lb:0,ub:50`
}
