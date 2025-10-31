package ies

// SCGFailureInformation-r12-criticalExtensions ::= CHOICE
const (
	ScgfailureinformationR12CriticalextensionsChoiceNothing = iota
	ScgfailureinformationR12CriticalextensionsChoiceC1
	ScgfailureinformationR12CriticalextensionsChoiceCriticalextensionsfuture
)

type ScgfailureinformationR12Criticalextensions struct {
	Choice                   uint64
	C1                       *ScgfailureinformationR12CriticalextensionsC1
	Criticalextensionsfuture *ScgfailureinformationR12CriticalextensionsCriticalextensionsfuture
}
