package ies

import "rrc/utils"

// NPRACH-ParametersList-NB-r13 ::= SEQUENCE OF NPRACH-Parameters-NB-r13
// SIZE (1.. maxNPRACH-Resources-NB-r13)
type NprachParameterslistNbR13 struct {
	Value utils.Sequence[NprachParametersNbR13]
}
