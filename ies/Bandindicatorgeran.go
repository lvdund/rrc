package ies

import "rrc/utils"

// BandIndicatorGERAN ::= ENUMERATED
type Bandindicatorgeran struct {
	Value utils.ENUMERATED
}

const (
	BandindicatorgeranDcs1800 = 0
	BandindicatorgeranPcs1900 = 1
)
