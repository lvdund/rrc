package ies

import "rrc/utils"

// GWUS-GroupsForServiceList-NB-r16 ::= SEQUENCE OF utils.INTEGER // SIZE (1..maxGWUS-ProbThresholds-NB-r16)
type GwusGroupsforservicelistNbR16 struct {
	Value []utils.INTEGER `lb:1,ub:maxGWUSProbthresholdsNbR16`
}
