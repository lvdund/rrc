package ies

// RRCRelease-criticalExtensions ::= CHOICE
const (
	RrcreleaseCriticalextensionsChoiceNothing = iota
	RrcreleaseCriticalextensionsChoiceRrcrelease
	RrcreleaseCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcreleaseCriticalextensions struct {
	Choice                   uint64
	Rrcrelease               *Rrcrelease
	Criticalextensionsfuture *RrcreleaseCriticalextensionsCriticalextensionsfuture
}
