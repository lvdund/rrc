package ies

// PRACH-ParametersListCE-r13 ::= SEQUENCE OF PRACH-ParametersCE-r13
// SIZE (1..maxCE-Level-r13)
type PrachParameterslistceR13 struct {
	Value []PrachParametersceR13 `lb:1,ub:maxCELevelR13`
}
