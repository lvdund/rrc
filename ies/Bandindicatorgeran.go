package ies

import "rrc/utils"

// BandIndicatorGERAN ::= ENUMERATED
type Bandindicatorgeran struct {
	Value utils.ENUMERATED
}

const (
	BandindicatorgeranEnumeratedNothing = iota
	BandindicatorgeranEnumeratedDcs1800
	BandindicatorgeranEnumeratedPcs1900
)
