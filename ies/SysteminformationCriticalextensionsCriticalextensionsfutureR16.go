package ies

// SystemInformation-criticalExtensions-criticalExtensionsFuture-r16 ::= CHOICE
const (
	SysteminformationCriticalextensionsCriticalextensionsfutureR16ChoiceNothing = iota
	SysteminformationCriticalextensionsCriticalextensionsfutureR16ChoicePossysteminformationR16
	SysteminformationCriticalextensionsCriticalextensionsfutureR16ChoiceCriticalextensionsfuture
)

type SysteminformationCriticalextensionsCriticalextensionsfutureR16 struct {
	Choice                   uint64
	PossysteminformationR16  *PossysteminformationR16
	Criticalextensionsfuture *SysteminformationCriticalextensionsCriticalextensionsfutureR16Criticalextensionsfuture
}
