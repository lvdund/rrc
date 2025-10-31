package ies

import "rrc/utils"

// PowSav-ParametersFRX-Diff-r16-maxMIMO-LayerPreference-r16 ::= ENUMERATED
type PowsavParametersfrxDiffR16MaxmimoLayerpreferenceR16 struct {
	Value utils.ENUMERATED
}

const (
	PowsavParametersfrxDiffR16MaxmimoLayerpreferenceR16EnumeratedNothing = iota
	PowsavParametersfrxDiffR16MaxmimoLayerpreferenceR16EnumeratedSupported
)
