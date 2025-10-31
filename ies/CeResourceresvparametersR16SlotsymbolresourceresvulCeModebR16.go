package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-slotSymbolResourceResvUL-CE-ModeB-r16 ::= ENUMERATED
type CeResourceresvparametersR16SlotsymbolresourceresvulCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SlotsymbolresourceresvulCeModebR16EnumeratedNothing = iota
	CeResourceresvparametersR16SlotsymbolresourceresvulCeModebR16EnumeratedSupported
)
