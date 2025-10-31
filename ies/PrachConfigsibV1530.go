package ies

// PRACH-ConfigSIB-v1530 ::= SEQUENCE
type PrachConfigsibV1530 struct {
	EdtPrachParameterslistceR15 []EdtPrachParametersceR15 `lb:1,ub:maxCELevelR13`
}
