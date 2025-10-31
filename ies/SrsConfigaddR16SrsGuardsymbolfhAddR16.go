package ies

import "rrc/utils"

// SRS-ConfigAdd-r16-srs-GuardSymbolFH-Add-r16 ::= ENUMERATED
type SrsConfigaddR16SrsGuardsymbolfhAddR16 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigaddR16SrsGuardsymbolfhAddR16EnumeratedNothing = iota
	SrsConfigaddR16SrsGuardsymbolfhAddR16EnumeratedEnabled
)
