package ies

// MeasurementReportAppLayerList-r17 ::= SEQUENCE OF MeasReportAppLayer-r17
// SIZE (1..maxNrofAppLayerMeas-r17)
type MeasurementreportapplayerlistR17 struct {
	Value []MeasreportapplayerR17 `lb:1,ub:maxNrofAppLayerMeasR17`
}
