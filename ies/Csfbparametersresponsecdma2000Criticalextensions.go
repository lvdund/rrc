package ies

// CSFBParametersResponseCDMA2000-criticalExtensions ::= CHOICE
const (
	Csfbparametersresponsecdma2000CriticalextensionsChoiceNothing = iota
	Csfbparametersresponsecdma2000CriticalextensionsChoiceCsfbparametersresponsecdma2000R8
	Csfbparametersresponsecdma2000CriticalextensionsChoiceCriticalextensionsfuture
)

type Csfbparametersresponsecdma2000Criticalextensions struct {
	Choice                           uint64
	Csfbparametersresponsecdma2000R8 *Csfbparametersresponsecdma2000R8
	Criticalextensionsfuture         *Csfbparametersresponsecdma2000CriticalextensionsCriticalextensionsfuture
}
