package ies

// MeasResultEUTRA ::= SEQUENCE
// Extensible
type Measresulteutra struct {
	EutraPhyscellid Physcellid
	Measresult      Measquantityresultseutra
	CgiInfo         *CgiInfoeutra
}
