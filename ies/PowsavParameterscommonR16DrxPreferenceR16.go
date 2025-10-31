package ies

import "rrc/utils"

// PowSav-ParametersCommon-r16-drx-Preference-r16 ::= ENUMERATED
type PowsavParameterscommonR16DrxPreferenceR16 struct {
	Value utils.ENUMERATED
}

const (
	PowsavParameterscommonR16DrxPreferenceR16EnumeratedNothing = iota
	PowsavParameterscommonR16DrxPreferenceR16EnumeratedSupported
)
