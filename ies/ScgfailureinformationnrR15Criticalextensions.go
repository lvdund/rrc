package ies

// SCGFailureInformationNR-r15-criticalExtensions ::= CHOICE
const (
	ScgfailureinformationnrR15CriticalextensionsChoiceNothing = iota
	ScgfailureinformationnrR15CriticalextensionsChoiceC1
	ScgfailureinformationnrR15CriticalextensionsChoiceCriticalextensionsfuture
)

type ScgfailureinformationnrR15Criticalextensions struct {
	Choice                   uint64
	C1                       *ScgfailureinformationnrR15CriticalextensionsC1
	Criticalextensionsfuture *ScgfailureinformationnrR15CriticalextensionsCriticalextensionsfuture
}
