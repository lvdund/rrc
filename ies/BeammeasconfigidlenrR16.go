package ies

import "rrc/utils"

// BeamMeasConfigIdleNR-r16 ::= SEQUENCE
type BeammeasconfigidlenrR16 struct {
	ReportquantityrsIndexnrR16 BeammeasconfigidlenrR16ReportquantityrsIndexnrR16
	MaxreportrsIndexR16        utils.INTEGER `lb:0,ub:maxRSIndexreportR15`
	ReportrsIndexresultsnrR16  utils.BOOLEAN
}
