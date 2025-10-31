package ies

import "rrc/utils"

// SystemInformationBlockType20-r13 ::= SEQUENCE
// Extensible
type Systeminformationblocktype20R13 struct {
	ScMcchRepetitionperiodR13   Systeminformationblocktype20R13ScMcchRepetitionperiodR13
	ScMcchOffsetR13             utils.INTEGER  `lb:0,ub:10`
	ScMcchFirstsubframeR13      utils.INTEGER  `lb:0,ub:9`
	ScMcchDurationR13           *utils.INTEGER `lb:0,ub:9`
	ScMcchModificationperiodR13 Systeminformationblocktype20R13ScMcchModificationperiodR13
	Latenoncriticalextension    *utils.OCTETSTRING
}
