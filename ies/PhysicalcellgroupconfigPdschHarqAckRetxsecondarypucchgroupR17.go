package ies

import "rrc/utils"

// PhysicalCellGroupConfig-pdsch-HARQ-ACK-RetxSecondaryPUCCHgroup-r17 ::= ENUMERATED
type PhysicalcellgroupconfigPdschHarqAckRetxsecondarypucchgroupR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPdschHarqAckRetxsecondarypucchgroupR17EnumeratedNothing = iota
	PhysicalcellgroupconfigPdschHarqAckRetxsecondarypucchgroupR17EnumeratedEnabled
)
