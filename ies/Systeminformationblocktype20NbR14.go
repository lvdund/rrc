package ies

import "rrc/utils"

// SystemInformationBlockType20-NB-r14 ::= SEQUENCE
// Extensible
type Systeminformationblocktype20NbR14 struct {
	NpdcchScMcchConfigR14       NpdcchScMcchConfigNbR14
	ScMcchCarrierconfigR14      Systeminformationblocktype20NbR14ScMcchCarrierconfigR14
	ScMcchRepetitionperiodR14   Systeminformationblocktype20NbR14ScMcchRepetitionperiodR14
	ScMcchOffsetR14             utils.INTEGER `lb:0,ub:10`
	ScMcchModificationperiodR14 Systeminformationblocktype20NbR14ScMcchModificationperiodR14
	ScMcchSchedulinginfoR14     *ScMcchSchedulinginfoNbR14
	Latenoncriticalextension    *utils.OCTETSTRING
}
