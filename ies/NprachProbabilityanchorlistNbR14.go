package ies

// NPRACH-ProbabilityAnchorList-NB-r14 ::= SEQUENCE OF NPRACH-ProbabilityAnchor-NB-r14
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type NprachProbabilityanchorlistNbR14 struct {
	Value []NprachProbabilityanchorNbR14 `lb:1,ub:maxNPRACHResourcesNbR13`
}
