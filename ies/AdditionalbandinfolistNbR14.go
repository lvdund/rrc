package ies

import "rrc/utils"

// AdditionalBandInfoList-NB-r14 ::= SEQUENCE OF FreqBandIndicator-NB-r13
// SIZE (1..maxMultiBands)
type AdditionalbandinfolistNbR14 struct {
	Value utils.Sequence[FreqbandindicatorNbR13]
}
