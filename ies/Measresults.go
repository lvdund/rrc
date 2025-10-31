package ies

// MeasResults ::= SEQUENCE
// Extensible
type Measresults struct {
	Measid               Measid
	Measresultpcell      MeasresultsMeasresultpcell
	Measresultneighcells *MeasresultsMeasresultneighcells
}
