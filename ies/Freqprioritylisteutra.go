package ies

import "rrc/utils"

// FreqPriorityListEUTRA ::= SEQUENCE OF FreqPriorityEUTRA
// SIZE (1..maxFreq)
type Freqprioritylisteutra struct {
	Value []Freqpriorityeutra
}
