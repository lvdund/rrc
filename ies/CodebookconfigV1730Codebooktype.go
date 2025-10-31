package ies

// CodebookConfig-v1730-codebookType ::= CHOICE
const (
	CodebookconfigV1730CodebooktypeChoiceNothing = iota
	CodebookconfigV1730CodebooktypeChoiceType1
)

type CodebookconfigV1730Codebooktype struct {
	Choice uint64
	Type1  *CodebookconfigV1730CodebooktypeType1 `lb:1,ub:2`
}
