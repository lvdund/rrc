package ies

import "rrc/utils"

// SystemInformationBlockType20-r13 ::= SEQUENCE
// Extensible
type Systeminformationblocktype20R13 struct {
	ScMcchRepetitionperiodR13   utils.ENUMERATED
	ScMcchOffsetR13             utils.INTEGER
	ScMcchFirstsubframeR13      utils.INTEGER
	ScMcchDurationR13           *utils.INTEGER
	ScMcchModificationperiodR13 utils.ENUMERATED
	Latenoncriticalextension    *utils.OCTETSTRING
}
