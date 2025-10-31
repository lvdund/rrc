package ies

// NPRACH-ParametersList-NB-r13 ::= SEQUENCE OF NPRACH-Parameters-NB-r13
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type NprachParameterslistNbR13 struct {
	Value []NprachParametersNbR13 `lb:1,ub:maxNPRACHResourcesNbR13`
}
