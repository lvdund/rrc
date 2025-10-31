package ies

import "rrc/utils"

// EUTRA-ParametersCommon-multiNS-Pmax-EUTRA ::= ENUMERATED
type EutraParameterscommonMultinsPmaxEutra struct {
	Value utils.ENUMERATED
}

const (
	EutraParameterscommonMultinsPmaxEutraEnumeratedNothing = iota
	EutraParameterscommonMultinsPmaxEutraEnumeratedSupported
)
