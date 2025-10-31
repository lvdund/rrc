package ies

// MeasConfig-s-MeasureConfig ::= CHOICE
const (
	MeasconfigSMeasureconfigChoiceNothing = iota
	MeasconfigSMeasureconfigChoiceSsbRsrp
	MeasconfigSMeasureconfigChoiceCsiRsrp
)

type MeasconfigSMeasureconfig struct {
	Choice  uint64
	SsbRsrp *RsrpRange
	CsiRsrp *RsrpRange
}
