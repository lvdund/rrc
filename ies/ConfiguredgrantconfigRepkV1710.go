package ies

import "rrc/utils"

// ConfiguredGrantConfig-repK-v1710 ::= ENUMERATED
type ConfiguredgrantconfigRepkV1710 struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigRepkV1710EnumeratedNothing = iota
	ConfiguredgrantconfigRepkV1710EnumeratedN12
	ConfiguredgrantconfigRepkV1710EnumeratedN16
	ConfiguredgrantconfigRepkV1710EnumeratedN24
	ConfiguredgrantconfigRepkV1710EnumeratedN32
)
