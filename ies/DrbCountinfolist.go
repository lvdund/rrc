package ies

// DRB-CountInfoList ::= SEQUENCE OF DRB-CountInfo
// SIZE (0..maxDRB)
type DrbCountinfolist struct {
	Value []DrbCountinfo `lb:0,ub:maxDRB`
}
