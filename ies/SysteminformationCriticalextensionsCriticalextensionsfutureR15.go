package ies

// SystemInformation-criticalExtensions-criticalExtensionsFuture-r15 ::= CHOICE
const (
	SysteminformationCriticalextensionsCriticalextensionsfutureR15ChoiceNothing = iota
	SysteminformationCriticalextensionsCriticalextensionsfutureR15ChoicePossysteminformationR15
	SysteminformationCriticalextensionsCriticalextensionsfutureR15ChoiceCriticalextensionsfuture
)

type SysteminformationCriticalextensionsCriticalextensionsfutureR15 struct {
	Choice                   uint64
	PossysteminformationR15  *PossysteminformationR15
	Criticalextensionsfuture *SysteminformationCriticalextensionsCriticalextensionsfutureR15Criticalextensionsfuture
}
