package ies

import "rrc/utils"

// PhysicalCellGroupConfig-simultaneousPUCCH-PUSCH-SecondaryPUCCHgroup-r17 ::= ENUMERATED
type PhysicalcellgroupconfigSimultaneouspucchPuschSecondarypucchgroupR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigSimultaneouspucchPuschSecondarypucchgroupR17EnumeratedNothing = iota
	PhysicalcellgroupconfigSimultaneouspucchPuschSecondarypucchgroupR17EnumeratedEnabled
)
