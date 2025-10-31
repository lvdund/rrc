package ies

// SystemInformation-criticalExtensions ::= CHOICE
const (
	SysteminformationCriticalextensionsChoiceNothing = iota
	SysteminformationCriticalextensionsChoiceSysteminformationR8
	SysteminformationCriticalextensionsChoiceCriticalextensionsfutureR15
)

type SysteminformationCriticalextensions struct {
	Choice                      uint64
	SysteminformationR8         *SysteminformationR8
	CriticalextensionsfutureR15 *SysteminformationCriticalextensionsCriticalextensionsfutureR15
}
