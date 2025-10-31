package ies

// UEPagingCoverageInformation-criticalExtensions-c1 ::= CHOICE
const (
	UepagingcoverageinformationCriticalextensionsC1ChoiceNothing = iota
	UepagingcoverageinformationCriticalextensionsC1ChoiceUepagingcoverageinformationR13
	UepagingcoverageinformationCriticalextensionsC1ChoiceSpare7
	UepagingcoverageinformationCriticalextensionsC1ChoiceSpare6
	UepagingcoverageinformationCriticalextensionsC1ChoiceSpare5
	UepagingcoverageinformationCriticalextensionsC1ChoiceSpare4
	UepagingcoverageinformationCriticalextensionsC1ChoiceSpare3
	UepagingcoverageinformationCriticalextensionsC1ChoiceSpare2
	UepagingcoverageinformationCriticalextensionsC1ChoiceSpare1
)

type UepagingcoverageinformationCriticalextensionsC1 struct {
	Choice                         uint64
	UepagingcoverageinformationR13 *UepagingcoverageinformationR13
	Spare7                         *struct{}
	Spare6                         *struct{}
	Spare5                         *struct{}
	Spare4                         *struct{}
	Spare3                         *struct{}
	Spare2                         *struct{}
	Spare1                         *struct{}
}
