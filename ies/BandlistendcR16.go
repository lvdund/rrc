package ies

import "rrc/utils"

// BandListENDC-r16 ::= SEQUENCE OF FreqBandIndicatorNR-r15
// SIZE (1.. maxBandsENDC-r16)
type BandlistendcR16 struct {
	Value utils.Sequence[FreqbandindicatornrR15]
}
