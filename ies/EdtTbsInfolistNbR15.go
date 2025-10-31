package ies

// EDT-TBS-InfoList-NB-r15 ::= SEQUENCE OF EDT-TBS-NB-r15
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type EdtTbsInfolistNbR15 struct {
	Value []EdtTbsNbR15 `lb:1,ub:maxNPRACHResourcesNbR13`
}
