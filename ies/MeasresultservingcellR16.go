package ies

// MeasResultServingCell-r16 ::= SEQUENCE
type MeasresultservingcellR16 struct {
	ResultsssbCell Measquantityresults
	Resultsssb     *MeasresultservingcellR16Resultsssb `lb:1,ub:maxNrofSSBsR16`
}
