package ies

import "rrc/utils"

// SRS-ConfigAp-v1310 ::= SEQUENCE
type SrsConfigapV1310 struct {
	TransmissioncombapV1310 *utils.INTEGER `lb:0,ub:3`
	CyclicshiftapV1310      *SrsConfigapV1310CyclicshiftapV1310
	TransmissioncombnumR13  *SrsConfigapV1310TransmissioncombnumR13
}
