package ies

import "rrc/utils"

// EUTRA-ParametersCommon-mfbi-EUTRA ::= ENUMERATED
type EutraParameterscommonMfbiEutra struct {
	Value utils.ENUMERATED
}

const (
	EutraParameterscommonMfbiEutraEnumeratedNothing = iota
	EutraParameterscommonMfbiEutraEnumeratedSupported
)
