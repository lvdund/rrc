package ies

import "rrc/utils"

// PhysicalCellGroupConfig-prioLowDG-HighCG-r17 ::= ENUMERATED
type PhysicalcellgroupconfigPriolowdgHighcgR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPriolowdgHighcgR17EnumeratedNothing = iota
	PhysicalcellgroupconfigPriolowdgHighcgR17EnumeratedEnabled
)
