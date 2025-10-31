package ies

// SCGFailureInformation-r12-criticalExtensions-c1 ::= CHOICE
const (
	ScgfailureinformationR12CriticalextensionsC1ChoiceNothing = iota
	ScgfailureinformationR12CriticalextensionsC1ChoiceScgfailureinformationR12
	ScgfailureinformationR12CriticalextensionsC1ChoiceSpare3
	ScgfailureinformationR12CriticalextensionsC1ChoiceSpare2
	ScgfailureinformationR12CriticalextensionsC1ChoiceSpare1
)

type ScgfailureinformationR12CriticalextensionsC1 struct {
	Choice                   uint64
	ScgfailureinformationR12 *ScgfailureinformationR12
	Spare3                   *struct{}
	Spare2                   *struct{}
	Spare1                   *struct{}
}
