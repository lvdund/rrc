package ies

// AdditionalReestabInfoList ::= SEQUENCE (SIZE (1..maxReestabInfo)) OF AdditionalReestabInfo
type Additionalreestabinfolist struct {
	Value []Additionalreestabinfo `lb:1,ub:maxReestabInfo`
}