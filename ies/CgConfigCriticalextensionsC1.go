package ies

// CG-Config-criticalExtensions-c1 ::= CHOICE
const (
	CgConfigCriticalextensionsC1ChoiceNothing = iota
	CgConfigCriticalextensionsC1ChoiceCgConfig
	CgConfigCriticalextensionsC1ChoiceSpare3
	CgConfigCriticalextensionsC1ChoiceSpare2
	CgConfigCriticalextensionsC1ChoiceSpare1
)

type CgConfigCriticalextensionsC1 struct {
	Choice   uint64
	CgConfig *CgConfig
	Spare3   *struct{}
	Spare2   *struct{}
	Spare1   *struct{}
}
