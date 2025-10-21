package ies

import "rrc/utils"

// InterFreqCarrierFreqList-NB-r13 ::= SEQUENCE OF InterFreqCarrierFreqInfo-NB-r13
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistNbR13 struct {
	Value utils.Sequence[InterfreqcarrierfreqinfoNbR13]
}
