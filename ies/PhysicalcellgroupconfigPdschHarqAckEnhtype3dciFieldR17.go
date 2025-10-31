package ies

import "rrc/utils"

// PhysicalCellGroupConfig-pdsch-HARQ-ACK-EnhType3DCI-Field-r17 ::= ENUMERATED
type PhysicalcellgroupconfigPdschHarqAckEnhtype3dciFieldR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPdschHarqAckEnhtype3dciFieldR17EnumeratedNothing = iota
	PhysicalcellgroupconfigPdschHarqAckEnhtype3dciFieldR17EnumeratedEnabled
)
