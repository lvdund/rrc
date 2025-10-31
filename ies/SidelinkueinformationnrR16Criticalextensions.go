package ies

// SidelinkUEInformationNR-r16-criticalExtensions ::= CHOICE
const (
	SidelinkueinformationnrR16CriticalextensionsChoiceNothing = iota
	SidelinkueinformationnrR16CriticalextensionsChoiceSidelinkueinformationnrR16
	SidelinkueinformationnrR16CriticalextensionsChoiceCriticalextensionsfuture
)

type SidelinkueinformationnrR16Criticalextensions struct {
	Choice                     uint64
	SidelinkueinformationnrR16 *SidelinkueinformationnrR16
	Criticalextensionsfuture   *SidelinkueinformationnrR16CriticalextensionsCriticalextensionsfuture
}
