package ies

import "rrc/utils"

// SRS-Resource-transmissionComb-n4 ::= SEQUENCE
type SrsResourceTransmissioncombN4 struct {
	ComboffsetN4  utils.INTEGER `lb:0,ub:3`
	CyclicshiftN4 utils.INTEGER `lb:0,ub:11`
}
