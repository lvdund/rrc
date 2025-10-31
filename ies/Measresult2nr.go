package ies

// MeasResult2NR ::= SEQUENCE
// Extensible
type Measresult2nr struct {
	Ssbfrequency              *ArfcnValuenr
	ReffreqcsiRs              *ArfcnValuenr
	Measresultservingcell     *Measresultnr
	Measresultneighcelllistnr *Measresultlistnr
}
