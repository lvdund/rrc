package ies

import "rrc/utils"

// PhysicalCellGroupConfig-bdFactorR-r16 ::= ENUMERATED
type PhysicalcellgroupconfigBdfactorrR16 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigBdfactorrR16EnumeratedNothing = iota
	PhysicalcellgroupconfigBdfactorrR16EnumeratedN1
)
