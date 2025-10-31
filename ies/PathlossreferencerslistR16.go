package ies

// PathlossReferenceRSList-r16 ::= SEQUENCE OF PathlossReferenceRS-r16
// SIZE (1..maxNrofSRS-PathlossReferenceRS-r16)
type PathlossreferencerslistR16 struct {
	Value []PathlossreferencersR16 `lb:1,ub:maxNrofSRSPathlossreferencersR16`
}
