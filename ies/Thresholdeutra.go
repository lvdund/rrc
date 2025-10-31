package ies

// ThresholdEUTRA ::= CHOICE
const (
	ThresholdeutraChoiceNothing = iota
	ThresholdeutraChoiceThresholdRsrp
	ThresholdeutraChoiceThresholdRsrq
)

type Thresholdeutra struct {
	Choice        uint64
	ThresholdRsrp *RsrpRange
	ThresholdRsrq *RsrqRange
}
