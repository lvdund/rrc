package ies

// ProximityIndication-r9-criticalExtensions-c1 ::= CHOICE
const (
	ProximityindicationR9CriticalextensionsC1ChoiceNothing = iota
	ProximityindicationR9CriticalextensionsC1ChoiceProximityindicationR9
	ProximityindicationR9CriticalextensionsC1ChoiceSpare3
	ProximityindicationR9CriticalextensionsC1ChoiceSpare2
	ProximityindicationR9CriticalextensionsC1ChoiceSpare1
)

type ProximityindicationR9CriticalextensionsC1 struct {
	Choice                uint64
	ProximityindicationR9 *ProximityindicationR9
	Spare3                *struct{}
	Spare2                *struct{}
	Spare1                *struct{}
}
