package ies

// MultiBandInfoList-NB-r13 ::= SEQUENCE OF MultiBandInfo-NB-r13
// SIZE (1..maxMultiBands)
type MultibandinfolistNbR13 struct {
	Value []MultibandinfoNbR13 `lb:1,ub:maxMultiBands`
}
