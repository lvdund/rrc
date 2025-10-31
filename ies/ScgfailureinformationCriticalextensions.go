package ies

// SCGFailureInformation-criticalExtensions ::= CHOICE
const (
	ScgfailureinformationCriticalextensionsChoiceNothing = iota
	ScgfailureinformationCriticalextensionsChoiceScgfailureinformation
	ScgfailureinformationCriticalextensionsChoiceCriticalextensionsfuture
)

type ScgfailureinformationCriticalextensions struct {
	Choice                   uint64
	Scgfailureinformation    *Scgfailureinformation
	Criticalextensionsfuture *ScgfailureinformationCriticalextensionsCriticalextensionsfuture
}
