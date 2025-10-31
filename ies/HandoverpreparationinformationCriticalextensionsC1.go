package ies

// HandoverPreparationInformation-criticalExtensions-c1 ::= CHOICE
const (
	HandoverpreparationinformationCriticalextensionsC1ChoiceNothing = iota
	HandoverpreparationinformationCriticalextensionsC1ChoiceHandoverpreparationinformation
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare3
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare2
	HandoverpreparationinformationCriticalextensionsC1ChoiceSpare1
)

type HandoverpreparationinformationCriticalextensionsC1 struct {
	Choice                         uint64
	Handoverpreparationinformation *Handoverpreparationinformation
	Spare3                         *struct{}
	Spare2                         *struct{}
	Spare1                         *struct{}
}
