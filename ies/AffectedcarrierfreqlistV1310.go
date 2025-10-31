package ies

// AffectedCarrierFreqList-v1310 ::= SEQUENCE OF AffectedCarrierFreq-v1310
// SIZE (1..maxFreqIDC-r11)
type AffectedcarrierfreqlistV1310 struct {
	Value []AffectedcarrierfreqV1310 `lb:1,ub:maxFreqIDCR11`
}
