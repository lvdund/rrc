package ies

import "rrc/utils"

// MAC-ParametersSidelinkCommon-r16-multipleConfiguredGrantsSidelink-r16 ::= ENUMERATED
type MacParameterssidelinkcommonR16MultipleconfiguredgrantssidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterssidelinkcommonR16MultipleconfiguredgrantssidelinkR16EnumeratedNothing = iota
	MacParameterssidelinkcommonR16MultipleconfiguredgrantssidelinkR16EnumeratedSupported
)
