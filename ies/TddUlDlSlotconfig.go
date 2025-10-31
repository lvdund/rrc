package ies

// TDD-UL-DL-SlotConfig ::= SEQUENCE
type TddUlDlSlotconfig struct {
	Slotindex TddUlDlSlotindex
	Symbols   *TddUlDlSlotconfigSymbols
}
