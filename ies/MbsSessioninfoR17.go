package ies

import "rrc/utils"

// MBS-SessionInfo-r17 ::= SEQUENCE
type MbsSessioninfoR17 struct {
	MbsSessionidR17              TmgiR17
	GRntiR17                     RntiValue
	MrbListbroadcastR17          MrbListbroadcastR17
	MtchSchedulinginfoR17        *DrxConfigptmIndexR17
	MtchNeighbourcellR17         *utils.BITSTRING `lb:maxNeighCellMBSR17,ub:maxNeighCellMBSR17`
	PdschConfigindexR17          *PdschConfigindexR17
	MtchSsbMappingwindowindexR17 *MtchSsbMappingwindowindexR17
}
