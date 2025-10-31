package ies

import "rrc/utils"

// MAC-CellGroupConfig-intraCG-Prioritization-r17 ::= ENUMERATED
type MacCellgroupconfigIntracgPrioritizationR17 struct {
	Value utils.ENUMERATED
}

const (
	MacCellgroupconfigIntracgPrioritizationR17EnumeratedNothing = iota
	MacCellgroupconfigIntracgPrioritizationR17EnumeratedEnabled
)
