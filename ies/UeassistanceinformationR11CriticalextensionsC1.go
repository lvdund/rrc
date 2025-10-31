package ies

// UEAssistanceInformation-r11-criticalExtensions-c1 ::= CHOICE
const (
	UeassistanceinformationR11CriticalextensionsC1ChoiceNothing = iota
	UeassistanceinformationR11CriticalextensionsC1ChoiceUeassistanceinformationR11
	UeassistanceinformationR11CriticalextensionsC1ChoiceSpare3
	UeassistanceinformationR11CriticalextensionsC1ChoiceSpare2
	UeassistanceinformationR11CriticalextensionsC1ChoiceSpare1
)

type UeassistanceinformationR11CriticalextensionsC1 struct {
	Choice                     uint64
	UeassistanceinformationR11 *UeassistanceinformationR11
	Spare3                     *struct{}
	Spare2                     *struct{}
	Spare1                     *struct{}
}
