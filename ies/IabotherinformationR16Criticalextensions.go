package ies

// IABOtherInformation-r16-criticalExtensions ::= CHOICE
const (
	IabotherinformationR16CriticalextensionsChoiceNothing = iota
	IabotherinformationR16CriticalextensionsChoiceIabotherinformationR16
	IabotherinformationR16CriticalextensionsChoiceCriticalextensionsfuture
)

type IabotherinformationR16Criticalextensions struct {
	Choice                   uint64
	IabotherinformationR16   *IabotherinformationR16
	Criticalextensionsfuture *IabotherinformationR16CriticalextensionsCriticalextensionsfuture
}
