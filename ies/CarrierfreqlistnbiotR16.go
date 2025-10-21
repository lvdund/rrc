package ies

import "rrc/utils"

// CarrierFreqListNBIOT-r16 ::= SEQUENCE OF CarrierFreqNBIOT-r16
// SIZE (1.. maxFreqNBIOT-r16)
type CarrierfreqlistnbiotR16 struct {
	Value utils.Sequence[CarrierfreqnbiotR16]
}
