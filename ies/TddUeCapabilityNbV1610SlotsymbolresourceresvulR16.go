package ies

import "rrc/utils"

// TDD-UE-Capability-NB-v1610-slotSymbolResourceResvUL-r16 ::= ENUMERATED
type TddUeCapabilityNbV1610SlotsymbolresourceresvulR16 struct {
	Value utils.ENUMERATED
}

const (
	TddUeCapabilityNbV1610SlotsymbolresourceresvulR16EnumeratedNothing = iota
	TddUeCapabilityNbV1610SlotsymbolresourceresvulR16EnumeratedSupported
)
