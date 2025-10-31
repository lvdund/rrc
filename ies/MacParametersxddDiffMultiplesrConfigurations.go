package ies

import "rrc/utils"

// MAC-ParametersXDD-Diff-multipleSR-Configurations ::= ENUMERATED
type MacParametersxddDiffMultiplesrConfigurations struct {
	Value utils.ENUMERATED
}

const (
	MacParametersxddDiffMultiplesrConfigurationsEnumeratedNothing = iota
	MacParametersxddDiffMultiplesrConfigurationsEnumeratedSupported
)
