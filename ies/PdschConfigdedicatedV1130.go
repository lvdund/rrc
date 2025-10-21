package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1130 ::= SEQUENCE
type PdschConfigdedicatedV1130 struct {
	DmrsConfigpdschR11                 *DmrsConfigR11
	QclOperation                       *utils.ENUMERATED
	ReMappingqclconfigtoreleaselistR11 *ReMappingqclconfigtoreleaselistR11
	ReMappingqclconfigtoaddmodlistR11  *ReMappingqclconfigtoaddmodlistR11
}
