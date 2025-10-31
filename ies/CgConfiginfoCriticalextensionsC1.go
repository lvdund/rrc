package ies

// CG-ConfigInfo-criticalExtensions-c1 ::= CHOICE
const (
	CgConfiginfoCriticalextensionsC1ChoiceNothing = iota
	CgConfiginfoCriticalextensionsC1ChoiceCgConfiginfo
	CgConfiginfoCriticalextensionsC1ChoiceSpare3
	CgConfiginfoCriticalextensionsC1ChoiceSpare2
	CgConfiginfoCriticalextensionsC1ChoiceSpare1
)

type CgConfiginfoCriticalextensionsC1 struct {
	Choice       uint64
	CgConfiginfo *CgConfiginfo
	Spare3       *struct{}
	Spare2       *struct{}
	Spare1       *struct{}
}
