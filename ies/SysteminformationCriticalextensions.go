package ies

// SystemInformation-criticalExtensions ::= CHOICE
const (
	SysteminformationCriticalextensionsChoiceNothing = iota
	SysteminformationCriticalextensionsChoiceSysteminformation
	SysteminformationCriticalextensionsChoiceCriticalextensionsfutureR16
)

type SysteminformationCriticalextensions struct {
	Choice                      uint64
	Systeminformation           *Systeminformation
	CriticalextensionsfutureR16 *SysteminformationCriticalextensionsCriticalextensionsfutureR16
}
