package ies

// RNReconfiguration-r10-criticalExtensions-c1 ::= CHOICE
const (
	RnreconfigurationR10CriticalextensionsC1ChoiceNothing = iota
	RnreconfigurationR10CriticalextensionsC1ChoiceRnreconfigurationR10
	RnreconfigurationR10CriticalextensionsC1ChoiceSpare3
	RnreconfigurationR10CriticalextensionsC1ChoiceSpare2
	RnreconfigurationR10CriticalextensionsC1ChoiceSpare1
)

type RnreconfigurationR10CriticalextensionsC1 struct {
	Choice               uint64
	RnreconfigurationR10 *RnreconfigurationR10
	Spare3               *struct{}
	Spare2               *struct{}
	Spare1               *struct{}
}
