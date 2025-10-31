package ies

// RACH-InfoList-NB-v1530 ::= SEQUENCE OF RACH-Info-NB-v1530
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type RachInfolistNbV1530 struct {
	Value []RachInfoNbV1530 `lb:1,ub:maxNPRACHResourcesNbR13`
}
