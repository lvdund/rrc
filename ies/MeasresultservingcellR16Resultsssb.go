package ies

import "rrc/utils"

// MeasResultServingCell-r16-resultsSSB ::= SEQUENCE
type MeasresultservingcellR16Resultsssb struct {
	BestSsbIndex    SsbIndex
	BestSsbResults  Measquantityresults
	Numberofgoodssb utils.INTEGER `lb:0,ub:maxNrofSSBsR16`
}
