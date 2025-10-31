package ies

import "rrc/utils"

// ConfiguredGrantConfig-resourceAllocation ::= ENUMERATED
type ConfiguredgrantconfigResourceallocation struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigResourceallocationEnumeratedNothing = iota
	ConfiguredgrantconfigResourceallocationEnumeratedResourceallocationtype0
	ConfiguredgrantconfigResourceallocationEnumeratedResourceallocationtype1
	ConfiguredgrantconfigResourceallocationEnumeratedDynamicswitch
)
