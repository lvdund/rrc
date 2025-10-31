package ies

// CodebookConfig-r17-codebookType ::= CHOICE
const (
	CodebookconfigR17CodebooktypeChoiceNothing = iota
	CodebookconfigR17CodebooktypeChoiceType1
	CodebookconfigR17CodebooktypeChoiceType2
)

type CodebookconfigR17Codebooktype struct {
	Choice uint64
	Type1  *CodebookconfigR17CodebooktypeType1
	Type2  *CodebookconfigR17CodebooktypeType2
}
