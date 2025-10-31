package ies

import "rrc/utils"

// BandNR-handoverUTRA-FDD-r16 ::= ENUMERATED
type BandnrHandoverutraFddR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrHandoverutraFddR16EnumeratedNothing = iota
	BandnrHandoverutraFddR16EnumeratedSupported
)
