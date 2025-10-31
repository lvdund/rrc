package ies

import "rrc/utils"

// ConfiguredGrantConfig-cg-UCI-Multiplexing-r16 ::= ENUMERATED
type ConfiguredgrantconfigCgUciMultiplexingR16 struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigCgUciMultiplexingR16EnumeratedNothing = iota
	ConfiguredgrantconfigCgUciMultiplexingR16EnumeratedEnabled
)
