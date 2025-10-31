package ies

// MeasResultUTRA ::= SEQUENCE
// Extensible
type Measresultutra struct {
	Physcellid MeasresultutraPhyscellid
	CgiInfo    *MeasresultutraCgiInfo
	Measresult *MeasresultutraMeasresult
}
