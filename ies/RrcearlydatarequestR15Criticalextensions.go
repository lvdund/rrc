package ies

// RRCEarlyDataRequest-r15-criticalExtensions ::= CHOICE
const (
	RrcearlydatarequestR15CriticalextensionsChoiceNothing = iota
	RrcearlydatarequestR15CriticalextensionsChoiceRrcearlydatarequestR15
	RrcearlydatarequestR15CriticalextensionsChoiceCriticalextensionsfuture
)

type RrcearlydatarequestR15Criticalextensions struct {
	Choice                   uint64
	RrcearlydatarequestR15   *RrcearlydatarequestR15
	Criticalextensionsfuture *RrcearlydatarequestR15CriticalextensionsCriticalextensionsfuture
}
