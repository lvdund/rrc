package ies

import "rrc/utils"

// ConfiguredGrantConfig-rbg-Size ::= ENUMERATED
type ConfiguredgrantconfigRbgSize struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigRbgSizeEnumeratedNothing = iota
	ConfiguredgrantconfigRbgSizeEnumeratedConfig2
)
