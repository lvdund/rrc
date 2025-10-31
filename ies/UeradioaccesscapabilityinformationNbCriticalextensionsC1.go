package ies

// UERadioAccessCapabilityInformation-NB-criticalExtensions-c1 ::= CHOICE
const (
	UeradioaccesscapabilityinformationNbCriticalextensionsC1ChoiceNothing = iota
	UeradioaccesscapabilityinformationNbCriticalextensionsC1ChoiceUeradioaccesscapabilityinformationR13
	UeradioaccesscapabilityinformationNbCriticalextensionsC1ChoiceSpare3
	UeradioaccesscapabilityinformationNbCriticalextensionsC1ChoiceSpare2
	UeradioaccesscapabilityinformationNbCriticalextensionsC1ChoiceSpare1
)

type UeradioaccesscapabilityinformationNbCriticalextensionsC1 struct {
	Choice                                uint64
	UeradioaccesscapabilityinformationR13 *UeradioaccesscapabilityinformationNb
	Spare3                                *struct{}
	Spare2                                *struct{}
	Spare1                                *struct{}
}
