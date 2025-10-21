package ies

import "rrc/utils"

// GWUS-GroupsForServiceList-r16 ::= SEQUENCE OF INTEGER
// SIZE (1..maxGWUS-ProbThresholds-r16)
type GwusGroupsforservicelistR16 struct {
	Value utils.Sequence[Integer]
}
