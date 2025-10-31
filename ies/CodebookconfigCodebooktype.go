package ies

// CodebookConfig-codebookType ::= CHOICE
const (
	CodebookconfigCodebooktypeChoiceNothing = iota
	CodebookconfigCodebooktypeChoiceType1
	CodebookconfigCodebooktypeChoiceType2
)

type CodebookconfigCodebooktype struct {
	Choice uint64
	Type1  *CodebookconfigCodebooktypeType1
	Type2  *CodebookconfigCodebooktypeType2
}
