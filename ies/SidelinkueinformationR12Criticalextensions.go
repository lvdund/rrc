package ies

// SidelinkUEInformation-r12-criticalExtensions ::= CHOICE
const (
	SidelinkueinformationR12CriticalextensionsChoiceNothing = iota
	SidelinkueinformationR12CriticalextensionsChoiceC1
	SidelinkueinformationR12CriticalextensionsChoiceCriticalextensionsfuture
)

type SidelinkueinformationR12Criticalextensions struct {
	Choice                   uint64
	C1                       *SidelinkueinformationR12CriticalextensionsC1
	Criticalextensionsfuture *SidelinkueinformationR12CriticalextensionsCriticalextensionsfuture
}
