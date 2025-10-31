package ies

// RRCReconfigurationComplete-criticalExtensions ::= CHOICE
const (
	RrcreconfigurationcompleteCriticalextensionsChoiceNothing = iota
	RrcreconfigurationcompleteCriticalextensionsChoiceRrcreconfigurationcomplete
	RrcreconfigurationcompleteCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcreconfigurationcompleteCriticalextensions struct {
	Choice                     uint64
	Rrcreconfigurationcomplete *Rrcreconfigurationcomplete
	Criticalextensionsfuture   *RrcreconfigurationcompleteCriticalextensionsCriticalextensionsfuture
}
