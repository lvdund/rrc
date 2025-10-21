package ies

import "rrc/utils"

// QuantityConfigUTRA ::= SEQUENCE
type Quantityconfigutra struct {
	MeasquantityutraFdd utils.ENUMERATED
	MeasquantityutraTdd utils.ENUMERATED
	Filtercoefficient   Filtercoefficient
}
