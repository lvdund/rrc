package ies

import "rrc/utils"

// NPRACH-ProbabilityAnchorList-NB-r14 ::= SEQUENCE OF NPRACH-ProbabilityAnchor-NB-r14
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type NprachProbabilityanchorlistNbR14 struct {
	Value []NprachProbabilityanchorNbR14
}
