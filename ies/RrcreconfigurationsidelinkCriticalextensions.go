package ies

// RRCReconfigurationSidelink-criticalExtensions ::= CHOICE
const (
	RrcreconfigurationsidelinkCriticalextensionsChoiceNothing = iota
	RrcreconfigurationsidelinkCriticalextensionsChoiceRrcreconfigurationsidelinkR16
	RrcreconfigurationsidelinkCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcreconfigurationsidelinkCriticalextensions struct {
	Choice                        uint64
	RrcreconfigurationsidelinkR16 *RrcreconfigurationsidelinkR16
	Criticalextensionsfuture      *RrcreconfigurationsidelinkCriticalextensionsCriticalextensionsfuture
}
