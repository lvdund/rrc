package ies

import "rrc/utils"

// InvalidSymbolPattern-r16-periodicityAndPattern-r16 ::= CHOICE
const (
	InvalidsymbolpatternR16PeriodicityandpatternR16ChoiceNothing = iota
	InvalidsymbolpatternR16PeriodicityandpatternR16ChoiceN2
	InvalidsymbolpatternR16PeriodicityandpatternR16ChoiceN4
	InvalidsymbolpatternR16PeriodicityandpatternR16ChoiceN5
	InvalidsymbolpatternR16PeriodicityandpatternR16ChoiceN8
	InvalidsymbolpatternR16PeriodicityandpatternR16ChoiceN10
	InvalidsymbolpatternR16PeriodicityandpatternR16ChoiceN20
	InvalidsymbolpatternR16PeriodicityandpatternR16ChoiceN40
)

type InvalidsymbolpatternR16PeriodicityandpatternR16 struct {
	Choice uint64
	N2     *utils.BITSTRING `lb:2,ub:2`
	N4     *utils.BITSTRING `lb:4,ub:4`
	N5     *utils.BITSTRING `lb:5,ub:5`
	N8     *utils.BITSTRING `lb:8,ub:8`
	N10    *utils.BITSTRING `lb:10,ub:10`
	N20    *utils.BITSTRING `lb:20,ub:20`
	N40    *utils.BITSTRING `lb:40,ub:40`
}
