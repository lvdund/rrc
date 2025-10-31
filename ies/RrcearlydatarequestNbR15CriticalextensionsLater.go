package ies

// RRCEarlyDataRequest-NB-r15-criticalExtensions-later ::= CHOICE
const (
	RrcearlydatarequestNbR15CriticalextensionsLaterChoiceNothing = iota
	RrcearlydatarequestNbR15CriticalextensionsLaterChoiceRrcearlydatarequestR16
	RrcearlydatarequestNbR15CriticalextensionsLaterChoiceCriticalextensionsfuture
)

type RrcearlydatarequestNbR15CriticalextensionsLater struct {
	Choice                   uint64
	RrcearlydatarequestR16   *Rrcearlydatarequest5gcNbR16
	Criticalextensionsfuture *RrcearlydatarequestNbR15CriticalextensionsLaterCriticalextensionsfuture
}
