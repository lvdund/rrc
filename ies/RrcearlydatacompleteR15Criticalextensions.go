package ies

// RRCEarlyDataComplete-r15-criticalExtensions ::= CHOICE
const (
	RrcearlydatacompleteR15CriticalextensionsChoiceNothing = iota
	RrcearlydatacompleteR15CriticalextensionsChoiceRrcearlydatacompleteR15
	RrcearlydatacompleteR15CriticalextensionsChoiceCriticalextensionsfuture
)

type RrcearlydatacompleteR15Criticalextensions struct {
	Choice                   uint64
	RrcearlydatacompleteR15  *RrcearlydatacompleteR15
	Criticalextensionsfuture *RrcearlydatacompleteR15CriticalextensionsCriticalextensionsfuture
}
