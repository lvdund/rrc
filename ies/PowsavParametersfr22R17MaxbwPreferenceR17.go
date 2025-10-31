package ies

import "rrc/utils"

// PowSav-ParametersFR2-2-r17-maxBW-Preference-r17 ::= ENUMERATED
type PowsavParametersfr22R17MaxbwPreferenceR17 struct {
	Value utils.ENUMERATED
}

const (
	PowsavParametersfr22R17MaxbwPreferenceR17EnumeratedNothing = iota
	PowsavParametersfr22R17MaxbwPreferenceR17EnumeratedSupported
)
