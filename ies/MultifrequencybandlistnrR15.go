package ies

import "rrc/utils"

// MultiFrequencyBandListNR-r15 ::= SEQUENCE OF FreqBandIndicatorNR-r15
// SIZE (1.. maxMultiBandsNR-r15)
type MultifrequencybandlistnrR15 struct {
	Value []FreqbandindicatornrR15
}
