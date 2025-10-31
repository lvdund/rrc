package ies

// UECapabilityInformation-criticalExtensions-c1 ::= CHOICE
const (
	UecapabilityinformationCriticalextensionsC1ChoiceNothing = iota
	UecapabilityinformationCriticalextensionsC1ChoiceUecapabilityinformationR8
	UecapabilityinformationCriticalextensionsC1ChoiceSpare7
	UecapabilityinformationCriticalextensionsC1ChoiceSpare6
	UecapabilityinformationCriticalextensionsC1ChoiceSpare5
	UecapabilityinformationCriticalextensionsC1ChoiceSpare4
	UecapabilityinformationCriticalextensionsC1ChoiceSpare3
	UecapabilityinformationCriticalextensionsC1ChoiceSpare2
	UecapabilityinformationCriticalextensionsC1ChoiceSpare1
)

type UecapabilityinformationCriticalextensionsC1 struct {
	Choice                    uint64
	UecapabilityinformationR8 *UecapabilityinformationR8
	Spare7                    *struct{}
	Spare6                    *struct{}
	Spare5                    *struct{}
	Spare4                    *struct{}
	Spare3                    *struct{}
	Spare2                    *struct{}
	Spare1                    *struct{}
}
