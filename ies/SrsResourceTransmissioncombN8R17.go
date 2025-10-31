package ies

import "rrc/utils"

// SRS-Resource-transmissionComb-n8-r17 ::= SEQUENCE
type SrsResourceTransmissioncombN8R17 struct {
	ComboffsetN8R17  utils.INTEGER `lb:0,ub:7`
	CyclicshiftN8R17 utils.INTEGER `lb:0,ub:5`
}
