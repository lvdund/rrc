package ies

import "rrc/utils"

// NPRACH-ParametersListFmt2-NB-r15 ::= SEQUENCE OF NPRACH-ParametersFmt2-NB-r15
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type NprachParameterslistfmt2NbR15 struct {
	Value []NprachParametersfmt2NbR15
}
