package ies

import "rrc/utils"

// PhysicalCellGroupConfig-twoQCLTypeDforPDCCHRepetition-r17 ::= ENUMERATED
type PhysicalcellgroupconfigTwoqcltypedforpdcchrepetitionR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigTwoqcltypedforpdcchrepetitionR17EnumeratedNothing = iota
	PhysicalcellgroupconfigTwoqcltypedforpdcchrepetitionR17EnumeratedEnabled
)
