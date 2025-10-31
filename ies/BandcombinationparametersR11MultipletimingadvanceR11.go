package ies

import "rrc/utils"

// BandCombinationParameters-r11-multipleTimingAdvance-r11 ::= ENUMERATED
type BandcombinationparametersR11MultipletimingadvanceR11 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersR11MultipletimingadvanceR11EnumeratedNothing = iota
	BandcombinationparametersR11MultipletimingadvanceR11EnumeratedSupported
)
