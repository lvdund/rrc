package ies

// MAC-ParametersSidelinkCommon-r16 ::= SEQUENCE
// Extensible
type MacParameterssidelinkcommonR16 struct {
	LcpRestrictionsidelinkR16           *MacParameterssidelinkcommonR16LcpRestrictionsidelinkR16
	MultipleconfiguredgrantssidelinkR16 *MacParameterssidelinkcommonR16MultipleconfiguredgrantssidelinkR16
	DrxOnsidelinkR17                    *MacParameterssidelinkcommonR16DrxOnsidelinkR17 `ext`
}
