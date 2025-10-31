package ies

// MeasResultEUTRA ::= SEQUENCE
// Extensible
type Measresulteutra struct {
	Physcellid Physcellid
	CgiInfo    *MeasresulteutraCgiInfo
	Measresult *MeasresulteutraMeasresult
}
