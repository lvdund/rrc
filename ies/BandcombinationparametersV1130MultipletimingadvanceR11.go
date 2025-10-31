package ies

import "rrc/utils"

// BandCombinationParameters-v1130-multipleTimingAdvance-r11 ::= ENUMERATED
type BandcombinationparametersV1130MultipletimingadvanceR11 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersV1130MultipletimingadvanceR11EnumeratedNothing = iota
	BandcombinationparametersV1130MultipletimingadvanceR11EnumeratedSupported
)
