package ies

// SCGFailureInformationNR-r15-criticalExtensions-c1 ::= CHOICE
const (
	ScgfailureinformationnrR15CriticalextensionsC1ChoiceNothing = iota
	ScgfailureinformationnrR15CriticalextensionsC1ChoiceScgfailureinformationnrR15
	ScgfailureinformationnrR15CriticalextensionsC1ChoiceSpare3
	ScgfailureinformationnrR15CriticalextensionsC1ChoiceSpare2
	ScgfailureinformationnrR15CriticalextensionsC1ChoiceSpare1
)

type ScgfailureinformationnrR15CriticalextensionsC1 struct {
	Choice                     uint64
	ScgfailureinformationnrR15 *ScgfailureinformationnrR15
	Spare3                     *struct{}
	Spare2                     *struct{}
	Spare1                     *struct{}
}
