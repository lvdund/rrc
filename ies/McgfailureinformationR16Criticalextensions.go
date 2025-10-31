package ies

// MCGFailureInformation-r16-criticalExtensions ::= CHOICE
const (
	McgfailureinformationR16CriticalextensionsChoiceNothing = iota
	McgfailureinformationR16CriticalextensionsChoiceMcgfailureinformationR16
	McgfailureinformationR16CriticalextensionsChoiceCriticalextensionsfuture
)

type McgfailureinformationR16Criticalextensions struct {
	Choice                   uint64
	McgfailureinformationR16 *McgfailureinformationR16
	Criticalextensionsfuture *McgfailureinformationR16CriticalextensionsCriticalextensionsfuture
}
