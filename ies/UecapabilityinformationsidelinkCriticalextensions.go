package ies

// UECapabilityInformationSidelink-criticalExtensions ::= CHOICE
const (
	UecapabilityinformationsidelinkCriticalextensionsChoiceNothing = iota
	UecapabilityinformationsidelinkCriticalextensionsChoiceUecapabilityinformationsidelinkR16
	UecapabilityinformationsidelinkCriticalextensionsChoiceCriticalextensionsfuture
)

type UecapabilityinformationsidelinkCriticalextensions struct {
	Choice                             uint64
	UecapabilityinformationsidelinkR16 *UecapabilityinformationsidelinkR16
	Criticalextensionsfuture           *UecapabilityinformationsidelinkCriticalextensionsCriticalextensionsfuture
}
