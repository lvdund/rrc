package ies

// CodebookConfig-r16-codebookType-type2-subType ::= CHOICE
const (
	CodebookconfigR16CodebooktypeType2SubtypeChoiceNothing = iota
	CodebookconfigR16CodebooktypeType2SubtypeChoiceTypeiiR16
	CodebookconfigR16CodebooktypeType2SubtypeChoiceTypeiiPortselectionR16
)

type CodebookconfigR16CodebooktypeType2Subtype struct {
	Choice                 uint64
	TypeiiR16              *CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16
	TypeiiPortselectionR16 *CodebookconfigR16CodebooktypeType2SubtypeTypeiiPortselectionR16
}
