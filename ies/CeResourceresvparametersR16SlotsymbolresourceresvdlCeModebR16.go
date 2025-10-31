package ies

import "rrc/utils"

// CE-ResourceResvParameters-r16-slotSymbolResourceResvDL-CE-ModeB-r16 ::= ENUMERATED
type CeResourceresvparametersR16SlotsymbolresourceresvdlCeModebR16 struct {
	Value utils.ENUMERATED
}

const (
	CeResourceresvparametersR16SlotsymbolresourceresvdlCeModebR16EnumeratedNothing = iota
	CeResourceresvparametersR16SlotsymbolresourceresvdlCeModebR16EnumeratedSupported
)
