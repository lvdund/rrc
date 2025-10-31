package ies

import "rrc/utils"

// SRS-PosResource-r16-transmissionComb-r16-n2-r16 ::= SEQUENCE
type SrsPosresourceR16TransmissioncombR16N2R16 struct {
	ComboffsetN2R16  utils.INTEGER `lb:0,ub:1`
	CyclicshiftN2R16 utils.INTEGER `lb:0,ub:7`
}
