package ies

import "rrc/utils"

// SchedulingRequestResourceConfigExt-v1610-phy-PriorityIndex-r16 ::= ENUMERATED
type SchedulingrequestresourceconfigextV1610PhyPriorityindexR16 struct {
	Value utils.ENUMERATED
}

const (
	SchedulingrequestresourceconfigextV1610PhyPriorityindexR16EnumeratedNothing = iota
	SchedulingrequestresourceconfigextV1610PhyPriorityindexR16EnumeratedP0
	SchedulingrequestresourceconfigextV1610PhyPriorityindexR16EnumeratedP1
)
