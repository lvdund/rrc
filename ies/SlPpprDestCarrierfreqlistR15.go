package ies

// SL-PPPR-Dest-CarrierFreqList-r15 ::= SEQUENCE OF SL-PPPR-Dest-CarrierFreq
// SIZE (1..maxSL-Dest-r12)
type SlPpprDestCarrierfreqlistR15 struct {
	Value []SlPpprDestCarrierfreq `lb:1,ub:maxSLDestR12`
}
