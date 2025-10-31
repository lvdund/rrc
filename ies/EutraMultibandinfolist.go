package ies

// EUTRA-MultiBandInfoList ::= SEQUENCE OF EUTRA-MultiBandInfo
// SIZE (1..maxMultiBands)
type EutraMultibandinfolist struct {
	Value []EutraMultibandinfo `lb:1,ub:maxMultiBands`
}
