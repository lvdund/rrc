package ies

// SCG-ConfigInfo-r12-criticalExtensions ::= CHOICE
const (
	ScgConfiginfoR12CriticalextensionsChoiceNothing = iota
	ScgConfiginfoR12CriticalextensionsChoiceC1
	ScgConfiginfoR12CriticalextensionsChoiceCriticalextensionsfuture
)

type ScgConfiginfoR12Criticalextensions struct {
	Choice                   uint64
	C1                       *ScgConfiginfoR12CriticalextensionsC1
	Criticalextensionsfuture *ScgConfiginfoR12CriticalextensionsCriticalextensionsfuture
}
