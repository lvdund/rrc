package ies

// RRCReestablishmentComplete-criticalExtensions ::= CHOICE
const (
	RrcreestablishmentcompleteCriticalextensionsChoiceNothing = iota
	RrcreestablishmentcompleteCriticalextensionsChoiceRrcreestablishmentcomplete
	RrcreestablishmentcompleteCriticalextensionsChoiceCriticalextensionsfuture
)

type RrcreestablishmentcompleteCriticalextensions struct {
	Choice                     uint64
	Rrcreestablishmentcomplete *Rrcreestablishmentcomplete
	Criticalextensionsfuture   *RrcreestablishmentcompleteCriticalextensionsCriticalextensionsfuture
}
