package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-slotSymbolResourceResvDL-CE-ModeA-r16 ::= ENUMERATED
type CeResourceresvparametersR16SlotsymbolresourceresvdlCeModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SlotsymbolresourceresvdlCeModeaR16EnumeratedNothing = iota
	CeResourceresvparametersR16SlotsymbolresourceresvdlCeModeaR16EnumeratedSupported
)
