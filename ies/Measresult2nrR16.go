package ies

// MeasResult2NR-r16 ::= SEQUENCE
type Measresult2nrR16 struct {
	SsbfrequencyR16   *ArfcnValuenr
	ReffreqcsiRsR16   *ArfcnValuenr
	MeasresultlistR16 Measresultlistnr
}
