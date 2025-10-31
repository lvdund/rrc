package ies

import "rrc/utils"

// BandCombinationParameters-r13-multipleTimingAdvance-r13 ::= ENUMERATED
type BandcombinationparametersR13MultipletimingadvanceR13 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersR13MultipletimingadvanceR13EnumeratedNothing = iota
	BandcombinationparametersR13MultipletimingadvanceR13EnumeratedSupported
)
