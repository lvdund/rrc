package ies

// SliceInfoList-r17 ::= SEQUENCE OF SliceInfo-r17
// SIZE (1..maxSliceInfo-r17)
type SliceinfolistR17 struct {
	Value []SliceinfoR17 `lb:1,ub:maxSliceInfoR17`
}
