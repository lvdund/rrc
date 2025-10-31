package ies

// MeasResultIdleEUTRA-r16 ::= SEQUENCE
// Extensible
type MeasresultidleeutraR16 struct {
	MeasresultspercarrierlistidleeutraR16 []MeasresultspercarrieridleeutraR16 `lb:1,ub:maxFreqIdleR16`
}
