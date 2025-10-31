package ies

import "rrc/utils"

// QuantityConfigGERAN-measQuantityGERAN ::= ENUMERATED
type QuantityconfiggeranMeasquantitygeran struct {
	Value utils.ENUMERATED
}

const (
	QuantityconfiggeranMeasquantitygeranEnumeratedNothing = iota
	QuantityconfiggeranMeasquantitygeranEnumeratedRssi
)
