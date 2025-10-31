package ies

// RNReconfiguration-r10-criticalExtensions ::= CHOICE
const (
	RnreconfigurationR10CriticalextensionsChoiceNothing = iota
	RnreconfigurationR10CriticalextensionsChoiceC1
	RnreconfigurationR10CriticalextensionsChoiceCriticalextensionsfuture
)

type RnreconfigurationR10Criticalextensions struct {
	Choice                   uint64
	C1                       *RnreconfigurationR10CriticalextensionsC1
	Criticalextensionsfuture *RnreconfigurationR10CriticalextensionsCriticalextensionsfuture
}
