package ies

import "rrc/utils"

// MBMS-InterFreqCarrierTypeList-r14 ::= SEQUENCE OF MBMS-CarrierType-r14
// SIZE (1..maxFreq)
type MbmsInterfreqcarriertypelistR14 struct {
	Value utils.Sequence[MbmsCarriertypeR14]
}
