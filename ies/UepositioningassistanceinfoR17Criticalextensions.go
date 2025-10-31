package ies

// UEPositioningAssistanceInfo-r17-criticalExtensions ::= CHOICE
const (
	UepositioningassistanceinfoR17CriticalextensionsChoiceNothing = iota
	UepositioningassistanceinfoR17CriticalextensionsChoiceUepositioningassistanceinfoR17
	UepositioningassistanceinfoR17CriticalextensionsChoiceCriticalextensionsfuture
)

type UepositioningassistanceinfoR17Criticalextensions struct {
	Choice                         uint64
	UepositioningassistanceinfoR17 *UepositioningassistanceinfoR17
	Criticalextensionsfuture       *UepositioningassistanceinfoR17CriticalextensionsCriticalextensionsfuture
}
