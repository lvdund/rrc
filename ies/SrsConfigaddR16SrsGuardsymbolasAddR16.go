package ies

import "rrc/utils"

// SRS-ConfigAdd-r16-srs-GuardSymbolAS-Add-r16 ::= ENUMERATED
type SrsConfigaddR16SrsGuardsymbolasAddR16 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigaddR16SrsGuardsymbolasAddR16EnumeratedNothing = iota
	SrsConfigaddR16SrsGuardsymbolasAddR16EnumeratedEnabled
)
