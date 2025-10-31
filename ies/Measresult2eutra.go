package ies

// MeasResult2EUTRA ::= SEQUENCE
// Extensible
type Measresult2eutra struct {
	Carrierfreq             ArfcnValueeutra
	Measresultservingcell   *Measresulteutra
	Measresultbestneighcell *Measresulteutra
}
