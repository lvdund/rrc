package ies

// MultiBandInfo-NB-r13 ::= SEQUENCE
type MultibandinfoNbR13 struct {
	FreqbandindicatorR13 *FreqbandindicatorNbR13
	FreqbandinfoR13      *NsPmaxlistNbR13
}
