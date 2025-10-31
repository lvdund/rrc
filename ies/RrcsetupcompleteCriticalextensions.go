package ies

// RRCSetupComplete-criticalExtensions ::= CHOICE
const (
	RrcsetupcompleteCriticalextensionsChoiceNothing = iota
	RrcsetupcompleteCriticalextensionsChoiceRrcsetupcomplete
	RrcsetupcompleteCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcsetupcompleteCriticalextensions struct {
	Choice                   uint64
	Rrcsetupcomplete         *Rrcsetupcomplete
	Criticalextensionsfuture *RrcsetupcompleteCriticalextensionsCriticalextensionsfuture
}
