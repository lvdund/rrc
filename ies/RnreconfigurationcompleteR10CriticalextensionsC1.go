package ies

// RNReconfigurationComplete-r10-criticalExtensions-c1 ::= CHOICE
const (
	RnreconfigurationcompleteR10CriticalextensionsC1ChoiceNothing = iota
	RnreconfigurationcompleteR10CriticalextensionsC1ChoiceRnreconfigurationcompleteR10
	RnreconfigurationcompleteR10CriticalextensionsC1ChoiceSpare3
	RnreconfigurationcompleteR10CriticalextensionsC1ChoiceSpare2
	RnreconfigurationcompleteR10CriticalextensionsC1ChoiceSpare1
)

type RnreconfigurationcompleteR10CriticalextensionsC1 struct {
	Choice                       uint64
	RnreconfigurationcompleteR10 *RnreconfigurationcompleteR10
	Spare3                       *struct{}
	Spare2                       *struct{}
	Spare1                       *struct{}
}
