package ies

// PathlossReferenceRSs-v1610 ::= SEQUENCE OF PUCCH-PathlossReferenceRS-r16
// SIZE (1..maxNrofPUCCH-PathlossReferenceRSsDiff-r16)
type PathlossreferencerssV1610 struct {
	Value []PucchPathlossreferencersR16 `lb:1,ub:maxNrofPUCCHPathlossreferencerssdiffR16`
}
