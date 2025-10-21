package ies

import "rrc/utils"

// InterFreqCarrierFreqList ::= SEQUENCE OF InterFreqCarrierFreqInfo
// SIZE (1..maxFreq)
type Interfreqcarrierfreqlist struct {
	Value []Interfreqcarrierfreqinfo
}
