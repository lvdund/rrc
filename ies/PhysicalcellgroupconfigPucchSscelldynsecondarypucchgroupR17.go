package ies

import "rrc/utils"

// PhysicalCellGroupConfig-pucch-sSCellDynSecondaryPUCCHgroup-r17 ::= ENUMERATED
type PhysicalcellgroupconfigPucchSscelldynsecondarypucchgroupR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPucchSscelldynsecondarypucchgroupR17EnumeratedNothing = iota
	PhysicalcellgroupconfigPucchSscelldynsecondarypucchgroupR17EnumeratedEnabled
)
