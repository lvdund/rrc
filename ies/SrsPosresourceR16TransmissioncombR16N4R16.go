package ies

import "rrc/utils"

// SRS-PosResource-r16-transmissionComb-r16-n4-r16 ::= SEQUENCE
type SrsPosresourceR16TransmissioncombR16N4R16 struct {
	ComboffsetN4R16  utils.INTEGER `lb:0,ub:3`
	CyclicshiftN4R16 utils.INTEGER `lb:0,ub:11`
}
