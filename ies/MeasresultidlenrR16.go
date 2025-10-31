package ies

// MeasResultIdleNR-r16 ::= SEQUENCE
// Extensible
type MeasresultidlenrR16 struct {
	MeasresultservingcellR16           *MeasresultidlenrR16MeasresultservingcellR16
	MeasresultspercarrierlistidlenrR16 *[]MeasresultspercarrieridlenrR16 `lb:1,ub:maxFreqIdleR16`
}
