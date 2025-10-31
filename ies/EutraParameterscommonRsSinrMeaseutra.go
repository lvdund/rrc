package ies

import "rrc/utils"

// EUTRA-ParametersCommon-rs-SINR-MeasEUTRA ::= ENUMERATED
type EutraParameterscommonRsSinrMeaseutra struct {
	Value utils.ENUMERATED
}

const (
	EutraParameterscommonRsSinrMeaseutraEnumeratedNothing = iota
	EutraParameterscommonRsSinrMeaseutraEnumeratedSupported
)
