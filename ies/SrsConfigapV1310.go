package ies

import "rrc/utils"

// SRS-ConfigAp-v1310 ::= SEQUENCE
type SrsConfigapV1310 struct {
	TransmissioncombapV1310 *utils.INTEGER
	CyclicshiftapV1310      *utils.ENUMERATED
	TransmissioncombnumR13  *utils.ENUMERATED
}
