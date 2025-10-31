package ies

import "rrc/utils"

// PowSav-ParametersFRX-Diff-r16-maxBW-Preference-r16 ::= ENUMERATED
type PowsavParametersfrxDiffR16MaxbwPreferenceR16 struct {
	Value utils.ENUMERATED
}

const (
	PowsavParametersfrxDiffR16MaxbwPreferenceR16EnumeratedNothing = iota
	PowsavParametersfrxDiffR16MaxbwPreferenceR16EnumeratedSupported
)
