package ies

// HandoverPreparationInformation-NB-criticalExtensions-c1 ::= CHOICE
const (
	HandoverpreparationinformationNbCriticalextensionsC1ChoiceNothing = iota
	HandoverpreparationinformationNbCriticalextensionsC1ChoiceHandoverpreparationinformationR13
	HandoverpreparationinformationNbCriticalextensionsC1ChoiceSpare3
	HandoverpreparationinformationNbCriticalextensionsC1ChoiceSpare2
	HandoverpreparationinformationNbCriticalextensionsC1ChoiceSpare1
)

type HandoverpreparationinformationNbCriticalextensionsC1 struct {
	Choice                            uint64
	HandoverpreparationinformationR13 *HandoverpreparationinformationNb
	Spare3                            *struct{}
	Spare2                            *struct{}
	Spare1                            *struct{}
}
