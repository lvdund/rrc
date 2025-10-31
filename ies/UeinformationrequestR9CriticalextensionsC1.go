package ies

// UEInformationRequest-r9-criticalExtensions-c1 ::= CHOICE
const (
	UeinformationrequestR9CriticalextensionsC1ChoiceNothing = iota
	UeinformationrequestR9CriticalextensionsC1ChoiceUeinformationrequestR9
	UeinformationrequestR9CriticalextensionsC1ChoiceSpare3
	UeinformationrequestR9CriticalextensionsC1ChoiceSpare2
	UeinformationrequestR9CriticalextensionsC1ChoiceSpare1
)

type UeinformationrequestR9CriticalextensionsC1 struct {
	Choice                 uint64
	UeinformationrequestR9 *UeinformationrequestR9
	Spare3                 *struct{}
	Spare2                 *struct{}
	Spare1                 *struct{}
}
