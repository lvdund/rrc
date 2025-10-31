package ies

import "rrc/utils"

// SC-MTCH-Info-NB-r14 ::= SEQUENCE
// Extensible
type ScMtchInfoNbR14 struct {
	ScMtchCarrierconfigR14        ScMtchInfoNbR14ScMtchCarrierconfigR14
	MbmssessioninfoR14            MbmssessioninfoR13
	GRntiR14                      utils.BITSTRING `lb:16,ub:16`
	ScMtchSchedulinginfoR14       *ScMtchSchedulinginfoNbR14
	ScMtchNeighbourcellR14        *utils.BITSTRING `lb:maxNeighCellScptmNbR14,ub:maxNeighCellScptmNbR14`
	NpdcchNpdschMaxtbsScMtchR14   ScMtchInfoNbR14NpdcchNpdschMaxtbsScMtchR14
	NpdcchNumrepetitionsScMtchR14 ScMtchInfoNbR14NpdcchNumrepetitionsScMtchR14
	NpdcchStartsfScMtchR14        ScMtchInfoNbR14NpdcchStartsfScMtchR14
	NpdcchOffsetScMtchR14         ScMtchInfoNbR14NpdcchOffsetScMtchR14
}
