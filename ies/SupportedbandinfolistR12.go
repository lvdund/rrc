package ies

// SupportedBandInfoList-r12 ::= SEQUENCE OF SupportedBandInfo-r12
// SIZE (1..maxBands)
type SupportedbandinfolistR12 struct {
	Value []SupportedbandinfoR12 `lb:1,ub:maxBands`
}
