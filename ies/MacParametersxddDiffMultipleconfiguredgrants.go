package ies

import "rrc/utils"

// MAC-ParametersXDD-Diff-multipleConfiguredGrants ::= ENUMERATED
type MacParametersxddDiffMultipleconfiguredgrants struct {
	Value utils.ENUMERATED
}

const (
	MacParametersxddDiffMultipleconfiguredgrantsEnumeratedNothing = iota
	MacParametersxddDiffMultipleconfiguredgrantsEnumeratedSupported
)
