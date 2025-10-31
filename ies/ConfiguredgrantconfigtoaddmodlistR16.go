package ies

// ConfiguredGrantConfigToAddModList-r16 ::= SEQUENCE OF ConfiguredGrantConfig
// SIZE (1..maxNrofConfiguredGrantConfig-r16)
type ConfiguredgrantconfigtoaddmodlistR16 struct {
	Value []Configuredgrantconfig `lb:1,ub:maxNrofConfiguredGrantConfigR16`
}
