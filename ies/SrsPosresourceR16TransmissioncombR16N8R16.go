package ies

import "rrc/utils"

// SRS-PosResource-r16-transmissionComb-r16-n8-r16 ::= SEQUENCE
type SrsPosresourceR16TransmissioncombR16N8R16 struct {
	ComboffsetN8R16  utils.INTEGER `lb:0,ub:7`
	CyclicshiftN8R16 utils.INTEGER `lb:0,ub:5`
}
