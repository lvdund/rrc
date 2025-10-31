package ies

// MultiBandInfoList-v10j0 ::= SEQUENCE OF NS-PmaxList-r10
// SIZE (1..maxMultiBands)
type MultibandinfolistV10j0 struct {
	Value []NsPmaxlistR10 `lb:1,ub:maxMultiBands`
}
