package ies

import "rrc/utils"

// PhysicalCellGroupConfig-pdsch-HARQ-ACK-Retx-r17 ::= ENUMERATED
type PhysicalcellgroupconfigPdschHarqAckRetxR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPdschHarqAckRetxR17EnumeratedNothing = iota
	PhysicalcellgroupconfigPdschHarqAckRetxR17EnumeratedEnabled
)
