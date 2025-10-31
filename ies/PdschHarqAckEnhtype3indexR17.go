package ies

import "rrc/utils"

// PDSCH-HARQ-ACK-EnhType3Index-r17 ::= utils.INTEGER (0..maxNrofEnhType3HARQ-ACK-1-r17)
type PdschHarqAckEnhtype3indexR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofEnhType3HARQAck1R17`
}
