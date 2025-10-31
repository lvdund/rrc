package ies

// RACH-InfoList-NB-r13 ::= SEQUENCE OF RACH-Info-NB-r13
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type RachInfolistNbR13 struct {
	Value []RachInfoNbR13 `lb:1,ub:maxNPRACHResourcesNbR13`
}
