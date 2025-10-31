package ies

import "rrc/utils"

// PhysicalCellGroupConfig-uci-MuxWithDiffPrio-r17 ::= ENUMERATED
type PhysicalcellgroupconfigUciMuxwithdiffprioR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigUciMuxwithdiffprioR17EnumeratedNothing = iota
	PhysicalcellgroupconfigUciMuxwithdiffprioR17EnumeratedEnabled
)
