package ies

// UEInformationResponse-r9-criticalExtensions-c1 ::= CHOICE
const (
	UeinformationresponseR9CriticalextensionsC1ChoiceNothing = iota
	UeinformationresponseR9CriticalextensionsC1ChoiceUeinformationresponseR9
	UeinformationresponseR9CriticalextensionsC1ChoiceSpare3
	UeinformationresponseR9CriticalextensionsC1ChoiceSpare2
	UeinformationresponseR9CriticalextensionsC1ChoiceSpare1
)

type UeinformationresponseR9CriticalextensionsC1 struct {
	Choice                  uint64
	UeinformationresponseR9 *UeinformationresponseR9
	Spare3                  *struct{}
	Spare2                  *struct{}
	Spare1                  *struct{}
}
