package ies

// VarMeasConfig-s-MeasureConfig ::= CHOICE
const (
	VarmeasconfigSMeasureconfigChoiceNothing = iota
	VarmeasconfigSMeasureconfigChoiceSsbRsrp
	VarmeasconfigSMeasureconfigChoiceCsiRsrp
)

type VarmeasconfigSMeasureconfig struct {
	Choice  uint64
	SsbRsrp *RsrpRange
	CsiRsrp *RsrpRange
}
