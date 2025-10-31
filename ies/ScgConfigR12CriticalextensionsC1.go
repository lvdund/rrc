package ies

// SCG-Config-r12-criticalExtensions-c1 ::= CHOICE
const (
	ScgConfigR12CriticalextensionsC1ChoiceNothing = iota
	ScgConfigR12CriticalextensionsC1ChoiceScgConfigR12
	ScgConfigR12CriticalextensionsC1ChoiceSpare7
	ScgConfigR12CriticalextensionsC1ChoiceSpare6
	ScgConfigR12CriticalextensionsC1ChoiceSpare5
	ScgConfigR12CriticalextensionsC1ChoiceSpare4
	ScgConfigR12CriticalextensionsC1ChoiceSpare3
	ScgConfigR12CriticalextensionsC1ChoiceSpare2
	ScgConfigR12CriticalextensionsC1ChoiceSpare1
)

type ScgConfigR12CriticalextensionsC1 struct {
	Choice       uint64
	ScgConfigR12 *ScgConfigR12
	Spare7       *struct{}
	Spare6       *struct{}
	Spare5       *struct{}
	Spare4       *struct{}
	Spare3       *struct{}
	Spare2       *struct{}
	Spare1       *struct{}
}
