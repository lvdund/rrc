package ies

import "rrc/utils"

// MultiBandInfoList ::= SEQUENCE OF FreqBandIndicator
// SIZE (1..maxMultiBands)
type Multibandinfolist struct {
	Value utils.Sequence[Freqbandindicator]
}
