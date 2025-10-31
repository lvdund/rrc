package ies

// MultiBandInfoList-v9e0 ::= SEQUENCE OF MultiBandInfo-v9e0
// SIZE (1..maxMultiBands)
type MultibandinfolistV9e0 struct {
	Value []MultibandinfoV9e0 `lb:1,ub:maxMultiBands`
}
