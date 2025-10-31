package ies

import "rrc/utils"

// PhysicalCellGroupConfig-nrdc-PCmode-FR2-r16 ::= ENUMERATED
type PhysicalcellgroupconfigNrdcPcmodeFr2R16 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigNrdcPcmodeFr2R16EnumeratedNothing = iota
	PhysicalcellgroupconfigNrdcPcmodeFr2R16EnumeratedSemi_Static_Mode1
	PhysicalcellgroupconfigNrdcPcmodeFr2R16EnumeratedSemi_Static_Mode2
	PhysicalcellgroupconfigNrdcPcmodeFr2R16EnumeratedDynamic
)
