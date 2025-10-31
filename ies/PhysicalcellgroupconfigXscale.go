package ies

import "rrc/utils"

// PhysicalCellGroupConfig-xScale ::= ENUMERATED
type PhysicalcellgroupconfigXscale struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigXscaleEnumeratedNothing = iota
	PhysicalcellgroupconfigXscaleEnumeratedDb0
	PhysicalcellgroupconfigXscaleEnumeratedDb6
	PhysicalcellgroupconfigXscaleEnumeratedSpare2
	PhysicalcellgroupconfigXscaleEnumeratedSpare1
)
