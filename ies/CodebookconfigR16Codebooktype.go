package ies

// CodebookConfig-r16-codebookType ::= CHOICE
const (
	CodebookconfigR16CodebooktypeChoiceNothing = iota
	CodebookconfigR16CodebooktypeChoiceType2
)

type CodebookconfigR16Codebooktype struct {
	Choice uint64
	Type2  *CodebookconfigR16CodebooktypeType2
}
