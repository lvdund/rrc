package ies

// MobilityFromNRCommand-criticalExtensions ::= CHOICE
const (
	MobilityfromnrcommandCriticalextensionsChoiceNothing = iota
	MobilityfromnrcommandCriticalextensionsChoiceMobilityfromnrcommand
	MobilityfromnrcommandCriticalextensionsChoiceCriticalextensionsfuture
)

type MobilityfromnrcommandCriticalextensions struct {
	Choice                   uint64
	Mobilityfromnrcommand    *Mobilityfromnrcommand
	Criticalextensionsfuture *MobilityfromnrcommandCriticalextensionsCriticalextensionsfuture
}
