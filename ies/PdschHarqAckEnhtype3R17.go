package ies

import "rrc/utils"

// PDSCH-HARQ-ACK-EnhType3-r17 ::= SEQUENCE
// Extensible
type PdschHarqAckEnhtype3R17 struct {
	PdschHarqAckEnhtype3indexR17 PdschHarqAckEnhtype3indexR17
	ApplicableR17                PdschHarqAckEnhtype3R17ApplicableR17
	PdschHarqAckEnhtype3ndiR17   *utils.BOOLEAN
	PdschHarqAckEnhtype3cbgR17   *utils.BOOLEAN
}
