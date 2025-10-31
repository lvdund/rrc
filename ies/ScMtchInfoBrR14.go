package ies

import "rrc/utils"

// SC-MTCH-Info-BR-r14 ::= SEQUENCE
// Extensible
type ScMtchInfoBrR14 struct {
	ScMtchCarrierfreqR14              ArfcnValueeutraR9
	MbmssessioninfoR14                MbmssessioninfoR13
	GRntiR14                          utils.BITSTRING `lb:16,ub:16`
	ScMtchSchedulinginfoR14           *ScMtchSchedulinginfoBrR14
	ScMtchNeighbourcellR14            *utils.BITSTRING `lb:maxNeighCellScptmR13,ub:maxNeighCellScptmR13`
	MpdcchNarrowbandScMtchR14         utils.INTEGER    `lb:0,ub:maxAvailNarrowBandsR13`
	MpdcchNumrepetitionScMtchR14      ScMtchInfoBrR14MpdcchNumrepetitionScMtchR14
	MpdcchStartsfScMtchR14            ScMtchInfoBrR14MpdcchStartsfScMtchR14
	MpdcchPdschHoppingconfigScMtchR14 ScMtchInfoBrR14MpdcchPdschHoppingconfigScMtchR14
	MpdcchPdschCemodeconfigScMtchR14  ScMtchInfoBrR14MpdcchPdschCemodeconfigScMtchR14
	MpdcchPdschMaxbandwidthScMtchR14  ScMtchInfoBrR14MpdcchPdschMaxbandwidthScMtchR14
	MpdcchOffsetScMtchR14             ScMtchInfoBrR14MpdcchOffsetScMtchR14
	PAR14                             *ScMtchInfoBrR14PAR14
}
