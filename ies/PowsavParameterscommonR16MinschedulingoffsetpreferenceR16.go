package ies

import "rrc/utils"

// PowSav-ParametersCommon-r16-minSchedulingOffsetPreference-r16 ::= ENUMERATED
type PowsavParameterscommonR16MinschedulingoffsetpreferenceR16 struct {
	Value utils.ENUMERATED
}

const (
	PowsavParameterscommonR16MinschedulingoffsetpreferenceR16EnumeratedNothing = iota
	PowsavParameterscommonR16MinschedulingoffsetpreferenceR16EnumeratedSupported
)
