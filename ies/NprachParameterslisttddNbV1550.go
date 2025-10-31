package ies

// NPRACH-ParametersListTDD-NB-v1550 ::= SEQUENCE OF NPRACH-ParametersTDD-NB-v1550
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type NprachParameterslisttddNbV1550 struct {
	Value []NprachParameterstddNbV1550 `lb:1,ub:maxNPRACHResourcesNbR13`
}
