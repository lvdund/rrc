package ies

import "rrc/utils"

// QuantityConfigWLAN-r13 ::= SEQUENCE
type QuantityconfigwlanR13 struct {
	MeasquantitywlanR13  utils.ENUMERATED
	FiltercoefficientR13 Filtercoefficient
}
