package ies

import "rrc/utils"

// ConfiguredGrantConfig-transformPrecoder ::= ENUMERATED
type ConfiguredgrantconfigTransformprecoder struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigTransformprecoderEnumeratedNothing = iota
	ConfiguredgrantconfigTransformprecoderEnumeratedEnabled
	ConfiguredgrantconfigTransformprecoderEnumeratedDisabled
)
