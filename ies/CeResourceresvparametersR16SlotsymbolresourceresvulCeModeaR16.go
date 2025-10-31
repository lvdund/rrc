package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-slotSymbolResourceResvUL-CE-ModeA-r16 ::= ENUMERATED
type CeResourceresvparametersR16SlotsymbolresourceresvulCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SlotsymbolresourceresvulCeModeaR16EnumeratedNothing = iota
	CeResourceresvparametersR16SlotsymbolresourceresvulCeModeaR16EnumeratedSupported
)
