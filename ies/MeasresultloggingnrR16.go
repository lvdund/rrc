package ies

import "rrc/utils"

// MeasResultLoggingNR-r16 ::= SEQUENCE
type MeasresultloggingnrR16 struct {
	PhyscellidR16      Physcellid
	ResultsssbCellR16  Measquantityresults
	NumberofgoodssbR16 *utils.INTEGER `lb:0,ub:maxNrofSSBsR16`
}
