package ies

import "rrc/utils"

// SL-BWP-Generic-r16-sl-StartSymbol-r16 ::= ENUMERATED
type SlBwpGenericR16SlStartsymbolR16 struct {
	Value utils.ENUMERATED
}

const (
	SlBwpGenericR16SlStartsymbolR16EnumeratedNothing = iota
	SlBwpGenericR16SlStartsymbolR16EnumeratedSym0
	SlBwpGenericR16SlStartsymbolR16EnumeratedSym1
	SlBwpGenericR16SlStartsymbolR16EnumeratedSym2
	SlBwpGenericR16SlStartsymbolR16EnumeratedSym3
	SlBwpGenericR16SlStartsymbolR16EnumeratedSym4
	SlBwpGenericR16SlStartsymbolR16EnumeratedSym5
	SlBwpGenericR16SlStartsymbolR16EnumeratedSym6
	SlBwpGenericR16SlStartsymbolR16EnumeratedSym7
)
