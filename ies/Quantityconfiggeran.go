package ies

import "rrc/utils"

// QuantityConfigGERAN ::= SEQUENCE
type Quantityconfiggeran struct {
	Measquantitygeran utils.ENUMERATED
	Filtercoefficient Filtercoefficient
}
