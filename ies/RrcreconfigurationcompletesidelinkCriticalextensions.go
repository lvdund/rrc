package ies

// RRCReconfigurationCompleteSidelink-criticalExtensions ::= CHOICE
const (
	RrcreconfigurationcompletesidelinkCriticalextensionsChoiceNothing = iota
	RrcreconfigurationcompletesidelinkCriticalextensionsChoiceRrcreconfigurationcompletesidelinkR16
	RrcreconfigurationcompletesidelinkCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcreconfigurationcompletesidelinkCriticalextensions struct {
	Choice                                uint64
	RrcreconfigurationcompletesidelinkR16 *RrcreconfigurationcompletesidelinkR16
	Criticalextensionsfuture              *RrcreconfigurationcompletesidelinkCriticalextensionsCriticalextensionsfuture
}
