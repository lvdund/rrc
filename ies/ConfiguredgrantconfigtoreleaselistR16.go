package ies

// ConfiguredGrantConfigToReleaseList-r16 ::= SEQUENCE OF ConfiguredGrantConfigIndex-r16
// SIZE (1..maxNrofConfiguredGrantConfig-r16)
type ConfiguredgrantconfigtoreleaselistR16 struct {
	Value []ConfiguredgrantconfigindexR16 `lb:1,ub:maxNrofConfiguredGrantConfigR16`
}
