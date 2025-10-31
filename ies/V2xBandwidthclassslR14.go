package ies

// V2X-BandwidthClassSL-r14 ::= SEQUENCE OF V2X-BandwidthClass-r14
// SIZE (1..maxBandwidthClass-r10)
type V2xBandwidthclassslR14 struct {
	Value []V2xBandwidthclassR14 `lb:1,ub:maxBandwidthClassR10`
}
