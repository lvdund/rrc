package ies

import "rrc/utils"

// ConfiguredGrantConfig-repK-RV ::= ENUMERATED
type ConfiguredgrantconfigRepkRv struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigRepkRvEnumeratedNothing = iota
	ConfiguredgrantconfigRepkRvEnumeratedS1_0231
	ConfiguredgrantconfigRepkRvEnumeratedS2_0303
	ConfiguredgrantconfigRepkRvEnumeratedS3_0000
)
