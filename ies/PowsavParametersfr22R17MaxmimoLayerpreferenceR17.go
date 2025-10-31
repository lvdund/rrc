package ies

import "rrc/utils"

// PowSav-ParametersFR2-2-r17-maxMIMO-LayerPreference-r17 ::= ENUMERATED
type PowsavParametersfr22R17MaxmimoLayerpreferenceR17 struct {
	Value utils.ENUMERATED
}

const (
	PowsavParametersfr22R17MaxmimoLayerpreferenceR17EnumeratedNothing = iota
	PowsavParametersfr22R17MaxmimoLayerpreferenceR17EnumeratedSupported
)
