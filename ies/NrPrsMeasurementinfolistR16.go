package ies

// NR-PRS-MeasurementInfoList-r16 ::= SEQUENCE OF NR-PRS-MeasurementInfo-r16
// SIZE (1..maxFreqLayers)
type NrPrsMeasurementinfolistR16 struct {
	Value []NrPrsMeasurementinfoR16 `lb:1,ub:maxFreqLayers`
}
