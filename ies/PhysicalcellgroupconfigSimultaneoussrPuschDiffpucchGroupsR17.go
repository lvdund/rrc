package ies

import "rrc/utils"

// PhysicalCellGroupConfig-simultaneousSR-PUSCH-diffPUCCH-Groups-r17 ::= ENUMERATED
type PhysicalcellgroupconfigSimultaneoussrPuschDiffpucchGroupsR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigSimultaneoussrPuschDiffpucchGroupsR17EnumeratedNothing = iota
	PhysicalcellgroupconfigSimultaneoussrPuschDiffpucchGroupsR17EnumeratedEnabled
)
