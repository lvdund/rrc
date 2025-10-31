package ies

// PDSCH-ConfigDedicated-v1130 ::= SEQUENCE
type PdschConfigdedicatedV1130 struct {
	DmrsConfigpdschR11                 *DmrsConfigR11
	QclOperation                       *PdschConfigdedicatedV1130QclOperation
	ReMappingqclconfigtoreleaselistR11 *ReMappingqclconfigtoreleaselistR11
	ReMappingqclconfigtoaddmodlistR11  *ReMappingqclconfigtoaddmodlistR11
}
