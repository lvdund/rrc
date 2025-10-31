package ies

import "rrc/utils"

// TDD-UL-DL-SlotConfig-symbols-explicit ::= SEQUENCE
type TddUlDlSlotconfigSymbolsExplicit struct {
	Nrofdownlinksymbols *utils.INTEGER `lb:0,ub:maxNrofSymbols1`
	Nrofuplinksymbols   *utils.INTEGER `lb:0,ub:maxNrofSymbols1`
}
