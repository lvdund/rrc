package ies

// UERadioAccessCapabilityInformation-criticalExtensions-c1 ::= CHOICE
const (
	UeradioaccesscapabilityinformationCriticalextensionsC1ChoiceNothing = iota
	UeradioaccesscapabilityinformationCriticalextensionsC1ChoiceUeradioaccesscapabilityinformationR8
	UeradioaccesscapabilityinformationCriticalextensionsC1ChoiceSpare7
	UeradioaccesscapabilityinformationCriticalextensionsC1ChoiceSpare6
	UeradioaccesscapabilityinformationCriticalextensionsC1ChoiceSpare5
	UeradioaccesscapabilityinformationCriticalextensionsC1ChoiceSpare4
	UeradioaccesscapabilityinformationCriticalextensionsC1ChoiceSpare3
	UeradioaccesscapabilityinformationCriticalextensionsC1ChoiceSpare2
	UeradioaccesscapabilityinformationCriticalextensionsC1ChoiceSpare1
)

type UeradioaccesscapabilityinformationCriticalextensionsC1 struct {
	Choice                               uint64
	UeradioaccesscapabilityinformationR8 *UeradioaccesscapabilityinformationR8
	Spare7                               *struct{}
	Spare6                               *struct{}
	Spare5                               *struct{}
	Spare4                               *struct{}
	Spare3                               *struct{}
	Spare2                               *struct{}
	Spare1                               *struct{}
}
