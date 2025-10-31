package ies

import "rrc/utils"

// MAC-ParametersSidelinkXDD-Diff-r16-multipleSR-ConfigurationsSidelink-r16 ::= ENUMERATED
type MacParameterssidelinkxddDiffR16MultiplesrConfigurationssidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterssidelinkxddDiffR16MultiplesrConfigurationssidelinkR16EnumeratedNothing = iota
	MacParameterssidelinkxddDiffR16MultiplesrConfigurationssidelinkR16EnumeratedSupported
)
