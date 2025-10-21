package ies

import "rrc/utils"

// QuantityConfigEUTRA ::= SEQUENCE
type Quantityconfigeutra struct {
	Filtercoefficientrsrp Filtercoefficient
	Filtercoefficientrsrq Filtercoefficient
}
