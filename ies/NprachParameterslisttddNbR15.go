package ies

import "rrc/utils"

// NPRACH-ParametersListTDD-NB-r15 ::= SEQUENCE OF NPRACH-ParametersTDD-NB-r15
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type NprachParameterslisttddNbR15 struct {
	Value utils.Sequence[NprachParameterstddNbR15]
}
