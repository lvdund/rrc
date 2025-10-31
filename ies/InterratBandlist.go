package ies

// InterRAT-BandList ::= SEQUENCE OF InterRAT-BandInfo
// SIZE (1..maxBands)
type InterratBandlist struct {
	Value []InterratBandinfo `lb:1,ub:maxBands`
}
