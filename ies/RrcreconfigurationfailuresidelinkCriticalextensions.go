package ies

// RRCReconfigurationFailureSidelink-criticalExtensions ::= CHOICE
const (
	RrcreconfigurationfailuresidelinkCriticalextensionsChoiceNothing = iota
	RrcreconfigurationfailuresidelinkCriticalextensionsChoiceRrcreconfigurationfailuresidelinkR16
	RrcreconfigurationfailuresidelinkCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcreconfigurationfailuresidelinkCriticalextensions struct {
	Choice                               uint64
	RrcreconfigurationfailuresidelinkR16 *RrcreconfigurationfailuresidelinkR16
	Criticalextensionsfuture             *RrcreconfigurationfailuresidelinkCriticalextensionsCriticalextensionsfuture
}
