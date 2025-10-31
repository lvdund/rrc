package ies

import "rrc/utils"

// VarMeasReportSL-r16 ::= SEQUENCE
type VarmeasreportslR16 struct {
	SlMeasidR16                 SlMeasidR16
	SlFrequencytriggeredlistR16 *[]ArfcnValuenr `lb:1,ub:maxNrofFreqSLR16`
	SlNumberofreportssentR16    utils.INTEGER
}
