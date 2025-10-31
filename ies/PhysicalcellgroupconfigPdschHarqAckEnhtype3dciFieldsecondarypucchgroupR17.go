package ies

import "rrc/utils"

// PhysicalCellGroupConfig-pdsch-HARQ-ACK-EnhType3DCI-FieldSecondaryPUCCHgroup-r17 ::= ENUMERATED
type PhysicalcellgroupconfigPdschHarqAckEnhtype3dciFieldsecondarypucchgroupR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPdschHarqAckEnhtype3dciFieldsecondarypucchgroupR17EnumeratedNothing = iota
	PhysicalcellgroupconfigPdschHarqAckEnhtype3dciFieldsecondarypucchgroupR17EnumeratedEnabled
)
