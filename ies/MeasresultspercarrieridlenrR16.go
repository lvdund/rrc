package ies

// MeasResultsPerCarrierIdleNR-r16 ::= SEQUENCE
// Extensible
type MeasresultspercarrieridlenrR16 struct {
	CarrierfreqR16                  ArfcnValuenr
	MeasresultspercelllistidlenrR16 []MeasresultspercellidlenrR16 `lb:1,ub:maxCellMeasIdleR16`
}
