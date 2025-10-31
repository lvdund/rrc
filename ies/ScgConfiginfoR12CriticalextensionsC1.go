package ies

// SCG-ConfigInfo-r12-criticalExtensions-c1 ::= CHOICE
const (
	ScgConfiginfoR12CriticalextensionsC1ChoiceNothing = iota
	ScgConfiginfoR12CriticalextensionsC1ChoiceScgConfiginfoR12
	ScgConfiginfoR12CriticalextensionsC1ChoiceSpare7
	ScgConfiginfoR12CriticalextensionsC1ChoiceSpare6
	ScgConfiginfoR12CriticalextensionsC1ChoiceSpare5
	ScgConfiginfoR12CriticalextensionsC1ChoiceSpare4
	ScgConfiginfoR12CriticalextensionsC1ChoiceSpare3
	ScgConfiginfoR12CriticalextensionsC1ChoiceSpare2
	ScgConfiginfoR12CriticalextensionsC1ChoiceSpare1
)

type ScgConfiginfoR12CriticalextensionsC1 struct {
	Choice           uint64
	ScgConfiginfoR12 *ScgConfiginfoR12
	Spare7           *struct{}
	Spare6           *struct{}
	Spare5           *struct{}
	Spare4           *struct{}
	Spare3           *struct{}
	Spare2           *struct{}
	Spare1           *struct{}
}
