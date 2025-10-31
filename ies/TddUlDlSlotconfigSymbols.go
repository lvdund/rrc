package ies

// TDD-UL-DL-SlotConfig-symbols ::= CHOICE
const (
	TddUlDlSlotconfigSymbolsChoiceNothing = iota
	TddUlDlSlotconfigSymbolsChoiceAlldownlink
	TddUlDlSlotconfigSymbolsChoiceAlluplink
	TddUlDlSlotconfigSymbolsChoiceExplicit
)

type TddUlDlSlotconfigSymbols struct {
	Choice      uint64
	Alldownlink *struct{}
	Alluplink   *struct{}
	Explicit    *TddUlDlSlotconfigSymbolsExplicit
}
