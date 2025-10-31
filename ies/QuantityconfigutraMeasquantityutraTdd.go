package ies

import "rrc/utils"

// QuantityConfigUTRA-measQuantityUTRA-TDD ::= ENUMERATED
type QuantityconfigutraMeasquantityutraTdd struct {
	Value utils.ENUMERATED
}

const (
	QuantityconfigutraMeasquantityutraTddEnumeratedNothing = iota
	QuantityconfigutraMeasquantityutraTddEnumeratedPccpch_Rscp
)
