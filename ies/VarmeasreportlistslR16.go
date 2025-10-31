package ies

// VarMeasReportListSL-r16 ::= SEQUENCE OF VarMeasReportSL-r16
// SIZE (1..maxNrofSL-MeasId-r16)
type VarmeasreportlistslR16 struct {
	Value []VarmeasreportslR16 `lb:1,ub:maxNrofSLMeasidR16`
}
