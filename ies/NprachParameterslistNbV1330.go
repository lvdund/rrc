package ies

// NPRACH-ParametersList-NB-v1330 ::= SEQUENCE OF NPRACH-Parameters-NB-v1330
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type NprachParameterslistNbV1330 struct {
	Value []NprachParametersNbV1330 `lb:1,ub:maxNPRACHResourcesNbR13`
}
