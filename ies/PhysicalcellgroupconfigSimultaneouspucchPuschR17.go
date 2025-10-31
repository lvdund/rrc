package ies

import "rrc/utils"

// PhysicalCellGroupConfig-simultaneousPUCCH-PUSCH-r17 ::= ENUMERATED
type PhysicalcellgroupconfigSimultaneouspucchPuschR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigSimultaneouspucchPuschR17EnumeratedNothing = iota
	PhysicalcellgroupconfigSimultaneouspucchPuschR17EnumeratedEnabled
)
