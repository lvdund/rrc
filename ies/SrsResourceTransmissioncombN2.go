package ies

import "rrc/utils"

// SRS-Resource-transmissionComb-n2 ::= SEQUENCE
type SrsResourceTransmissioncombN2 struct {
	ComboffsetN2  utils.INTEGER `lb:0,ub:1`
	CyclicshiftN2 utils.INTEGER `lb:0,ub:7`
}
