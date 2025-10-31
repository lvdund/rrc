package ies

// MeasResultServMO ::= SEQUENCE
// Extensible
type Measresultservmo struct {
	Servcellid              Servcellindex
	Measresultservingcell   Measresultnr
	Measresultbestneighcell *Measresultnr
}
