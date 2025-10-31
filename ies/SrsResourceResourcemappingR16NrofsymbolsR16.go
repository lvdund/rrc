package ies

import "rrc/utils"

// SRS-Resource-resourceMapping-r16-nrofSymbols-r16 ::= ENUMERATED
type SrsResourceResourcemappingR16NrofsymbolsR16 struct {
	Value utils.ENUMERATED
}

const (
	SrsResourceResourcemappingR16NrofsymbolsR16EnumeratedNothing = iota
	SrsResourceResourcemappingR16NrofsymbolsR16EnumeratedN1
	SrsResourceResourcemappingR16NrofsymbolsR16EnumeratedN2
	SrsResourceResourcemappingR16NrofsymbolsR16EnumeratedN4
)
