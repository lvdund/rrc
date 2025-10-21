package ies

import "rrc/utils"

// FreqPriorityListNR-r15 ::= SEQUENCE OF FreqPriorityNR-r15
// SIZE (1..maxFreq)
type FreqprioritylistnrR15 struct {
	Value utils.Sequence[FreqprioritynrR15]
}
