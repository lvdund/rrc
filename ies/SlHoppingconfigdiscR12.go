package ies

import "rrc/utils"

// SL-HoppingConfigDisc-r12 ::= SEQUENCE
type SlHoppingconfigdiscR12 struct {
	AR12 utils.INTEGER `lb:0,ub:200`
	BR12 utils.INTEGER `lb:0,ub:10`
	CR12 SlHoppingconfigdiscR12CR12
}
