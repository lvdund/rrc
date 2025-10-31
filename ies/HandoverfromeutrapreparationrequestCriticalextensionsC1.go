package ies

// HandoverFromEUTRAPreparationRequest-criticalExtensions-c1 ::= CHOICE
const (
	HandoverfromeutrapreparationrequestCriticalextensionsC1ChoiceNothing = iota
	HandoverfromeutrapreparationrequestCriticalextensionsC1ChoiceHandoverfromeutrapreparationrequestR8
	HandoverfromeutrapreparationrequestCriticalextensionsC1ChoiceSpare3
	HandoverfromeutrapreparationrequestCriticalextensionsC1ChoiceSpare2
	HandoverfromeutrapreparationrequestCriticalextensionsC1ChoiceSpare1
)

type HandoverfromeutrapreparationrequestCriticalextensionsC1 struct {
	Choice                                uint64
	HandoverfromeutrapreparationrequestR8 *HandoverfromeutrapreparationrequestR8
	Spare3                                *struct{}
	Spare2                                *struct{}
	Spare1                                *struct{}
}
