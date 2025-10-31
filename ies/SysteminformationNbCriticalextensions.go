package ies

// SystemInformation-NB-criticalExtensions ::= CHOICE
const (
	SysteminformationNbCriticalextensionsChoiceNothing = iota
	SysteminformationNbCriticalextensionsChoiceSysteminformationR13
	SysteminformationNbCriticalextensionsChoiceCriticalextensionsfuture
)

type SysteminformationNbCriticalextensions struct {
	Choice                   uint64
	SysteminformationR13     *SysteminformationNbR13
	Criticalextensionsfuture *SysteminformationNbCriticalextensionsCriticalextensionsfuture
}
