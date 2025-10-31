package ies

// UERadioPagingInformation-criticalExtensions-c1 ::= CHOICE
const (
	UeradiopaginginformationCriticalextensionsC1ChoiceNothing = iota
	UeradiopaginginformationCriticalextensionsC1ChoiceUeradiopaginginformation
	UeradiopaginginformationCriticalextensionsC1ChoiceSpare7
	UeradiopaginginformationCriticalextensionsC1ChoiceSpare6
	UeradiopaginginformationCriticalextensionsC1ChoiceSpare5
	UeradiopaginginformationCriticalextensionsC1ChoiceSpare4
	UeradiopaginginformationCriticalextensionsC1ChoiceSpare3
	UeradiopaginginformationCriticalextensionsC1ChoiceSpare2
	UeradiopaginginformationCriticalextensionsC1ChoiceSpare1
)

type UeradiopaginginformationCriticalextensionsC1 struct {
	Choice                   uint64
	Ueradiopaginginformation *Ueradiopaginginformation
	Spare7                   *struct{}
	Spare6                   *struct{}
	Spare5                   *struct{}
	Spare4                   *struct{}
	Spare3                   *struct{}
	Spare2                   *struct{}
	Spare1                   *struct{}
}
