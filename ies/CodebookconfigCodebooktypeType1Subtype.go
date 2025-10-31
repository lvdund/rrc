package ies

// CodebookConfig-codebookType-type1-subType ::= CHOICE
const (
	CodebookconfigCodebooktypeType1SubtypeChoiceNothing = iota
	CodebookconfigCodebooktypeType1SubtypeChoiceTypeiSinglepanel
	CodebookconfigCodebooktypeType1SubtypeChoiceTypeiMultipanel
)

type CodebookconfigCodebooktypeType1Subtype struct {
	Choice           uint64
	TypeiSinglepanel *CodebookconfigCodebooktypeType1SubtypeTypeiSinglepanel
	TypeiMultipanel  *CodebookconfigCodebooktypeType1SubtypeTypeiMultipanel
}
