package ies

// TDD-UL-DL-SlotConfig-IAB-MT-r16-symbols-IAB-MT-r16 ::= CHOICE
const (
	TddUlDlSlotconfigIabMtR16SymbolsIabMtR16ChoiceNothing = iota
	TddUlDlSlotconfigIabMtR16SymbolsIabMtR16ChoiceAlldownlinkR16
	TddUlDlSlotconfigIabMtR16SymbolsIabMtR16ChoiceAlluplinkR16
	TddUlDlSlotconfigIabMtR16SymbolsIabMtR16ChoiceExplicitR16
	TddUlDlSlotconfigIabMtR16SymbolsIabMtR16ChoiceExplicitIabMtR16
)

type TddUlDlSlotconfigIabMtR16SymbolsIabMtR16 struct {
	Choice           uint64
	AlldownlinkR16   *struct{}
	AlluplinkR16     *struct{}
	ExplicitR16      *TddUlDlSlotconfigIabMtR16SymbolsIabMtR16ExplicitR16
	ExplicitIabMtR16 *TddUlDlSlotconfigIabMtR16SymbolsIabMtR16ExplicitIabMtR16
}
