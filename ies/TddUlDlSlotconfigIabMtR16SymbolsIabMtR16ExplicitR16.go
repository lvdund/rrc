package ies

import "rrc/utils"

// TDD-UL-DL-SlotConfig-IAB-MT-r16-symbols-IAB-MT-r16-explicit-r16 ::= SEQUENCE
type TddUlDlSlotconfigIabMtR16SymbolsIabMtR16ExplicitR16 struct {
	NrofdownlinksymbolsR16 *utils.INTEGER `lb:0,ub:maxNrofSymbols1`
	NrofuplinksymbolsR16   *utils.INTEGER `lb:0,ub:maxNrofSymbols1`
}
