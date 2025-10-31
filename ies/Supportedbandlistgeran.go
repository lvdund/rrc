package ies

// SupportedBandListGERAN ::= SEQUENCE OF SupportedBandGERAN
// SIZE (1..maxBands)
type Supportedbandlistgeran struct {
	Value []Supportedbandgeran `lb:1,ub:maxBands`
}
