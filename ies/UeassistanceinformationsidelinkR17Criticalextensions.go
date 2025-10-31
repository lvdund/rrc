package ies

// UEAssistanceInformationSidelink-r17-criticalExtensions ::= CHOICE
const (
	UeassistanceinformationsidelinkR17CriticalextensionsChoiceNothing = iota
	UeassistanceinformationsidelinkR17CriticalextensionsChoiceUeassistanceinformationsidelinkR17
	UeassistanceinformationsidelinkR17CriticalextensionsChoiceCriticalextensionsfuture
)

type UeassistanceinformationsidelinkR17Criticalextensions struct {
	Choice                             uint64
	UeassistanceinformationsidelinkR17 *UeassistanceinformationsidelinkR17
	Criticalextensionsfuture           *UeassistanceinformationsidelinkR17CriticalextensionsCriticalextensionsfuture
}
