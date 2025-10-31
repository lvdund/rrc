package ies

import "rrc/utils"

// TDD-UE-Capability-NB-v1610-slotSymbolResourceResvDL-r16 ::= ENUMERATED
type TddUeCapabilityNbV1610SlotsymbolresourceresvdlR16 struct {
	Value utils.ENUMERATED
}

const (
	TddUeCapabilityNbV1610SlotsymbolresourceresvdlR16EnumeratedNothing = iota
	TddUeCapabilityNbV1610SlotsymbolresourceresvdlR16EnumeratedSupported
)
