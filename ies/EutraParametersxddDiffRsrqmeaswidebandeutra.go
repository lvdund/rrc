package ies

import "rrc/utils"

// EUTRA-ParametersXDD-Diff-rsrqMeasWidebandEUTRA ::= ENUMERATED
type EutraParametersxddDiffRsrqmeaswidebandeutra struct {
	Value utils.ENUMERATED
}

const (
	EutraParametersxddDiffRsrqmeaswidebandeutraEnumeratedNothing = iota
	EutraParametersxddDiffRsrqmeaswidebandeutraEnumeratedSupported
)
