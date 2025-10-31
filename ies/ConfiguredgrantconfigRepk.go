package ies

import "rrc/utils"

// ConfiguredGrantConfig-repK ::= ENUMERATED
type ConfiguredgrantconfigRepk struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigRepkEnumeratedNothing = iota
	ConfiguredgrantconfigRepkEnumeratedN1
	ConfiguredgrantconfigRepkEnumeratedN2
	ConfiguredgrantconfigRepkEnumeratedN4
	ConfiguredgrantconfigRepkEnumeratedN8
)
