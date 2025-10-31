package ies

import "rrc/utils"

// PDSCH-HARQ-ACK-EnhType3-r17-applicable-r17 ::= CHOICE
const (
	PdschHarqAckEnhtype3R17ApplicableR17ChoiceNothing = iota
	PdschHarqAckEnhtype3R17ApplicableR17ChoicePercc
	PdschHarqAckEnhtype3R17ApplicableR17ChoicePerharq
)

type PdschHarqAckEnhtype3R17ApplicableR17 struct {
	Choice  uint64
	Percc   *[]utils.INTEGER   `lb:1,ub:maxNrofServingCells`
	Perharq *[]utils.BITSTRING `lb:1,ub:maxNrofServingCells`
}
