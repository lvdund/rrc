package ies

import "rrc/utils"

// PhysicalCellGroupConfig-pucch-sSCellDyn-r17 ::= ENUMERATED
type PhysicalcellgroupconfigPucchSscelldynR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPucchSscelldynR17EnumeratedNothing = iota
	PhysicalcellgroupconfigPucchSscelldynR17EnumeratedEnabled
)
