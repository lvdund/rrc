package ies

import "rrc/utils"

// PhysicalCellGroupConfig-nrdc-PCmode-FR1-r16 ::= ENUMERATED
type PhysicalcellgroupconfigNrdcPcmodeFr1R16 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigNrdcPcmodeFr1R16EnumeratedNothing = iota
	PhysicalcellgroupconfigNrdcPcmodeFr1R16EnumeratedSemi_Static_Mode1
	PhysicalcellgroupconfigNrdcPcmodeFr1R16EnumeratedSemi_Static_Mode2
	PhysicalcellgroupconfigNrdcPcmodeFr1R16EnumeratedDynamic
)
