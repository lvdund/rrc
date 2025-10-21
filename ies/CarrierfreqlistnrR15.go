package ies

import "rrc/utils"

// CarrierFreqListNR-r15 ::= SEQUENCE OF CarrierFreqNR-r15
// SIZE (1..maxFreq)
type CarrierfreqlistnrR15 struct {
	Value utils.Sequence[CarrierfreqnrR15]
}
