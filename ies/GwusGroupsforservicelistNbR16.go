package ies

import "rrc/utils"

// GWUS-GroupsForServiceList-NB-r16 ::= SEQUENCE OF INTEGER
// SIZE (1..maxGWUS-ProbThresholds-NB-r16)
type GwusGroupsforservicelistNbR16 struct {
	Value utils.Sequence[Integer]
}
