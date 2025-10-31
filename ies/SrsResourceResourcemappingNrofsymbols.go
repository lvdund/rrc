package ies

import "rrc/utils"

// SRS-Resource-resourceMapping-nrofSymbols ::= ENUMERATED
type SrsResourceResourcemappingNrofsymbols struct {
	Value utils.ENUMERATED
}

const (
	SrsResourceResourcemappingNrofsymbolsEnumeratedNothing = iota
	SrsResourceResourcemappingNrofsymbolsEnumeratedN1
	SrsResourceResourcemappingNrofsymbolsEnumeratedN2
	SrsResourceResourcemappingNrofsymbolsEnumeratedN4
)
