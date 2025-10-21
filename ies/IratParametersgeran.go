package ies

import "rrc/utils"

// IRAT-ParametersGERAN ::= SEQUENCE
type IratParametersgeran struct {
	Supportedbandlistgeran Supportedbandlistgeran
	InterratPsHoTogeran    bool
}
