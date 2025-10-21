package ies

import "rrc/utils"

// NPRACH-ParametersList-NB-r14 ::= SEQUENCE OF NPRACH-Parameters-NB-r14
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type NprachParameterslistNbR14 struct {
	Value utils.Sequence[NprachParametersNbR14]
}
