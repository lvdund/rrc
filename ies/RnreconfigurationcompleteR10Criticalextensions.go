package ies

// RNReconfigurationComplete-r10-criticalExtensions ::= CHOICE
const (
	RnreconfigurationcompleteR10CriticalextensionsChoiceNothing = iota
	RnreconfigurationcompleteR10CriticalextensionsChoiceC1
	RnreconfigurationcompleteR10CriticalextensionsChoiceCriticalextensionsfuture
)

type RnreconfigurationcompleteR10Criticalextensions struct {
	Choice                   uint64
	C1                       *RnreconfigurationcompleteR10CriticalextensionsC1
	Criticalextensionsfuture *RnreconfigurationcompleteR10CriticalextensionsCriticalextensionsfuture
}
