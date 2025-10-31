package ies

// MCGFailureInformation-r16-criticalExtensions ::= CHOICE
const (
	McgfailureinformationR16CriticalextensionsChoiceNothing = iota
	McgfailureinformationR16CriticalextensionsChoiceMcgfailureinformation
	McgfailureinformationR16CriticalextensionsChoiceCriticalextensionsfuture
)

type McgfailureinformationR16Criticalextensions struct {
	Choice                   uint64
	Mcgfailureinformation    *McgfailureinformationR16
	Criticalextensionsfuture *McgfailureinformationR16CriticalextensionsCriticalextensionsfuture
}
