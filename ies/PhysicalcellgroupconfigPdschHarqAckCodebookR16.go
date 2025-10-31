package ies

import "rrc/utils"

// PhysicalCellGroupConfig-pdsch-HARQ-ACK-Codebook-r16 ::= ENUMERATED
type PhysicalcellgroupconfigPdschHarqAckCodebookR16 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPdschHarqAckCodebookR16EnumeratedNothing = iota
	PhysicalcellgroupconfigPdschHarqAckCodebookR16EnumeratedEnhanceddynamic
)
