package ies

// ProximityIndication-r9-criticalExtensions ::= CHOICE
const (
	ProximityindicationR9CriticalextensionsChoiceNothing = iota
	ProximityindicationR9CriticalextensionsChoiceC1
	ProximityindicationR9CriticalextensionsChoiceCriticalextensionsfuture
)

type ProximityindicationR9Criticalextensions struct {
	Choice                   uint64
	C1                       *ProximityindicationR9CriticalextensionsC1
	Criticalextensionsfuture *ProximityindicationR9CriticalextensionsCriticalextensionsfuture
}
