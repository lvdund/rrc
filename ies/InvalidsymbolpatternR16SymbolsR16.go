package ies

import "rrc/utils"

// InvalidSymbolPattern-r16-symbols-r16 ::= CHOICE
const (
	InvalidsymbolpatternR16SymbolsR16ChoiceNothing = iota
	InvalidsymbolpatternR16SymbolsR16ChoiceOneslot
	InvalidsymbolpatternR16SymbolsR16ChoiceTwoslots
)

type InvalidsymbolpatternR16SymbolsR16 struct {
	Choice   uint64
	Oneslot  *utils.BITSTRING `lb:14,ub:14`
	Twoslots *utils.BITSTRING `lb:28,ub:28`
}
