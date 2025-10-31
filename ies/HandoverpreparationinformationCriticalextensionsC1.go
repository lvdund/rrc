package ies

// HandoverPreparationInformation-criticalExtensions-c1 ::= CHOICE
const (
	HandoverpreparationinformationCriticalextensionsC1ChoiceNothing = iota
	HandoverpreparationinformationCriticalextensionsC1ChoiceHandoverpreparationinformationR8
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare7
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare6
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare5
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare4
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare3
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare2
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare1
)

type HandoverpreparationinformationCriticalextensionsC1 struct {
	Choice                           uint64
	HandoverpreparationinformationR8 *HandoverpreparationinformationR8
	Spare7                           *struct{}
	Spare6                           *struct{}
	Spare5                           *struct{}
	Spare4                           *struct{}
	Spare3                           *struct{}
	Spare2                           *struct{}
	Spare1                           *struct{}
}
