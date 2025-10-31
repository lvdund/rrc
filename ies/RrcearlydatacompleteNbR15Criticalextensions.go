package ies

// RRCEarlyDataComplete-NB-r15-criticalExtensions ::= CHOICE
const (
	RrcearlydatacompleteNbR15CriticalextensionsChoiceNothing = iota
	RrcearlydatacompleteNbR15CriticalextensionsChoiceRrcearlydatacompleteR15
	RrcearlydatacompleteNbR15CriticalextensionsChoiceCriticalextensionsfuture
)

type RrcearlydatacompleteNbR15Criticalextensions struct {
	Choice                   uint64
	RrcearlydatacompleteR15  *RrcearlydatacompleteNbR15
	Criticalextensionsfuture *RrcearlydatacompleteNbR15CriticalextensionsCriticalextensionsfuture
}
