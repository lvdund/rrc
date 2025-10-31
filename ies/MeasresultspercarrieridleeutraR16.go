package ies

// MeasResultsPerCarrierIdleEUTRA-r16 ::= SEQUENCE
// Extensible
type MeasresultspercarrieridleeutraR16 struct {
	CarrierfreqeutraR16                ArfcnValueeutra
	MeasresultspercelllistidleeutraR16 []MeasresultspercellidleeutraR16 `lb:1,ub:maxCellMeasIdleR16`
}
