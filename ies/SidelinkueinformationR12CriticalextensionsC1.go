package ies

// SidelinkUEInformation-r12-criticalExtensions-c1 ::= CHOICE
const (
	SidelinkueinformationR12CriticalextensionsC1ChoiceNothing = iota
	SidelinkueinformationR12CriticalextensionsC1ChoiceSidelinkueinformationR12
	SidelinkueinformationR12CriticalextensionsC1ChoiceSpare3
	SidelinkueinformationR12CriticalextensionsC1ChoiceSpare2
	SidelinkueinformationR12CriticalextensionsC1ChoiceSpare1
)

type SidelinkueinformationR12CriticalextensionsC1 struct {
	Choice                   uint64
	SidelinkueinformationR12 *SidelinkueinformationR12
	Spare3                   *struct{}
	Spare2                   *struct{}
	Spare1                   *struct{}
}
