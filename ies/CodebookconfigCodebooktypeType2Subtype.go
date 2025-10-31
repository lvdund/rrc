package ies

// CodebookConfig-codebookType-type2-subType ::= CHOICE
const (
	CodebookconfigCodebooktypeType2SubtypeChoiceNothing = iota
	CodebookconfigCodebooktypeType2SubtypeChoiceTypeii
	CodebookconfigCodebooktypeType2SubtypeChoiceTypeiiPortselection
)

type CodebookconfigCodebooktypeType2Subtype struct {
	Choice              uint64
	Typeii              *CodebookconfigCodebooktypeType2SubtypeTypeii
	TypeiiPortselection *CodebookconfigCodebooktypeType2SubtypeTypeiiPortselection
}
