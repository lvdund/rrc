package ies

import "rrc/utils"

// NZP-FrequencyDensity-r14 ::= ENUMERATED
type NzpFrequencydensityR14 struct {
	Value utils.ENUMERATED
}

const (
	NzpFrequencydensityR14EnumeratedNothing = iota
	NzpFrequencydensityR14EnumeratedD1
	NzpFrequencydensityR14EnumeratedD2
	NzpFrequencydensityR14EnumeratedD3
)
