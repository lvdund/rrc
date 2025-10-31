package ies

import "rrc/utils"

// PhysicalCellGroupConfig-uci-MuxWithDiffPrioSecondaryPUCCHgroup-r17 ::= ENUMERATED
type PhysicalcellgroupconfigUciMuxwithdiffpriosecondarypucchgroupR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigUciMuxwithdiffpriosecondarypucchgroupR17EnumeratedNothing = iota
	PhysicalcellgroupconfigUciMuxwithdiffpriosecondarypucchgroupR17EnumeratedEnabled
)
