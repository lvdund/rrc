package ies

import "rrc/utils"

// PowSav-ParametersCommon-r16-maxCC-Preference-r16 ::= ENUMERATED
type PowsavParameterscommonR16MaxccPreferenceR16 struct {
	Value utils.ENUMERATED
}

const (
	PowsavParameterscommonR16MaxccPreferenceR16EnumeratedNothing = iota
	PowsavParameterscommonR16MaxccPreferenceR16EnumeratedSupported
)
