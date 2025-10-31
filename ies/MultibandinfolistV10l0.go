package ies

// MultiBandInfoList-v10l0 ::= SEQUENCE OF NS-PmaxList-v10l0
// SIZE (1..maxMultiBands)
type MultibandinfolistV10l0 struct {
	Value []NsPmaxlistV10l0 `lb:1,ub:maxMultiBands`
}
