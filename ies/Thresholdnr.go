package ies

// ThresholdNR ::= SEQUENCE
type Thresholdnr struct {
	Thresholdrsrp *RsrpRange
	Thresholdrsrq *RsrqRange
	Thresholdsinr *SinrRange
}
