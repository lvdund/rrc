package ies

// SCG-Config-r12-criticalExtensions ::= CHOICE
const (
	ScgConfigR12CriticalextensionsChoiceNothing = iota
	ScgConfigR12CriticalextensionsChoiceC1
	ScgConfigR12CriticalextensionsChoiceCriticalextensionsfuture
)

type ScgConfigR12Criticalextensions struct {
	Choice                   uint64
	C1                       *ScgConfigR12CriticalextensionsC1
	Criticalextensionsfuture *ScgConfigR12CriticalextensionsCriticalextensionsfuture
}
