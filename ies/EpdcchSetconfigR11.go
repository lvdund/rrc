package ies

import "rrc/utils"

// EPDCCH-SetConfig-r11 ::= SEQUENCE
// Extensible
type EpdcchSetconfigR11 struct {
	SetconfigidR11               EpdcchSetconfigidR11
	TransmissiontypeR11          EpdcchSetconfigR11TransmissiontypeR11
	ResourceblockassignmentR11   EpdcchSetconfigR11ResourceblockassignmentR11
	DmrsScramblingsequenceintR11 utils.INTEGER `lb:0,ub:503`
	PucchResourcestartoffsetR11  utils.INTEGER `lb:0,ub:2047`
	ReMappingqclConfigidR11      *PdschReMappingqclConfigidR11
}
