package ies

// UERadioPagingInformation-NB-criticalExtensions-c1 ::= CHOICE
const (
	UeradiopaginginformationNbCriticalextensionsC1ChoiceNothing = iota
	UeradiopaginginformationNbCriticalextensionsC1ChoiceUeradiopaginginformationR13
	UeradiopaginginformationNbCriticalextensionsC1ChoiceSpare3
	UeradiopaginginformationNbCriticalextensionsC1ChoiceSpare2
	UeradiopaginginformationNbCriticalextensionsC1ChoiceSpare1
)

type UeradiopaginginformationNbCriticalextensionsC1 struct {
	Choice                      uint64
	UeradiopaginginformationR13 *UeradiopaginginformationNb
	Spare3                      *struct{}
	Spare2                      *struct{}
	Spare1                      *struct{}
}
