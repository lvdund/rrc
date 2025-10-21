package ies

import "rrc/utils"

// FreqBandIndicatorListEUTRA-r12 ::= SEQUENCE OF FreqBandIndicator-r11
// SIZE (1..maxBands)
type FreqbandindicatorlisteutraR12 struct {
	Value []FreqbandindicatorR11
}
