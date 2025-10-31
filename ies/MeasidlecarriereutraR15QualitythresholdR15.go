package ies

// MeasIdleCarrierEUTRA-r15-qualityThreshold-r15 ::= SEQUENCE
type MeasidlecarriereutraR15QualitythresholdR15 struct {
	IdlersrpThresholdR15 *RsrpRange
	IdlersrqThresholdR15 *RsrqRangeR13
}
