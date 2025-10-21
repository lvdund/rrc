package ies

import "rrc/utils"

// MIMO-CA-ParametersPerBoBC-r13 ::= SEQUENCE
type MimoCaParametersperbobcR13 struct {
	Parameterstm9R13  *MimoCaParametersperbobcpertmR13
	Parameterstm10R13 *MimoCaParametersperbobcpertmR13
}
