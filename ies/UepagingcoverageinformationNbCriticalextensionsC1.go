package ies

// UEPagingCoverageInformation-NB-criticalExtensions-c1 ::= CHOICE
const (
	UepagingcoverageinformationNbCriticalextensionsC1ChoiceNothing = iota
	UepagingcoverageinformationNbCriticalextensionsC1ChoiceUepagingcoverageinformationR13
	UepagingcoverageinformationNbCriticalextensionsC1ChoiceSpare3
	UepagingcoverageinformationNbCriticalextensionsC1ChoiceSpare2
	UepagingcoverageinformationNbCriticalextensionsC1ChoiceSpare1
)

type UepagingcoverageinformationNbCriticalextensionsC1 struct {
	Choice                         uint64
	UepagingcoverageinformationR13 *UepagingcoverageinformationNb
	Spare3                         *struct{}
	Spare2                         *struct{}
	Spare1                         *struct{}
}
