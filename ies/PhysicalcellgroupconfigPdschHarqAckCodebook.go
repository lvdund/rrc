package ies

import "rrc/utils"

// PhysicalCellGroupConfig-pdsch-HARQ-ACK-Codebook ::= ENUMERATED
type PhysicalcellgroupconfigPdschHarqAckCodebook struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigPdschHarqAckCodebookEnumeratedNothing = iota
	PhysicalcellgroupconfigPdschHarqAckCodebookEnumeratedSemistatic
	PhysicalcellgroupconfigPdschHarqAckCodebookEnumeratedDynamic
)
