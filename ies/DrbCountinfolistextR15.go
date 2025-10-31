package ies

// DRB-CountInfoListExt-r15 ::= SEQUENCE OF DRB-CountInfo
// SIZE (1..maxDRBExt-r15)
type DrbCountinfolistextR15 struct {
	Value []DrbCountinfo `lb:1,ub:maxDRBExtR15`
}
