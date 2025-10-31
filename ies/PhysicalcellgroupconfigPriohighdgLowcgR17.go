package ies

import "rrc/utils"

// PhysicalCellGroupConfig-prioHighDG-LowCG-r17 ::= ENUMERATED
type PhysicalcellgroupconfigPriohighdgLowcgR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPriohighdgLowcgR17EnumeratedNothing = iota
	PhysicalcellgroupconfigPriohighdgLowcgR17EnumeratedEnabled
)
