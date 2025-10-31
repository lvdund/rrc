package ies

// RRCReconfiguration-criticalExtensions ::= CHOICE
const (
	RrcreconfigurationCriticalextensionsChoiceNothing = iota
	RrcreconfigurationCriticalextensionsChoiceRrcreconfiguration
	RrcreconfigurationCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcreconfigurationCriticalextensions struct {
	Choice                   uint64
	Rrcreconfiguration       *Rrcreconfiguration
	Criticalextensionsfuture *RrcreconfigurationCriticalextensionsCriticalextensionsfuture
}
