package ies

// MobilityFromEUTRACommand-criticalExtensions ::= CHOICE
const (
	MobilityfromeutracommandCriticalextensionsChoiceNothing = iota
	MobilityfromeutracommandCriticalextensionsChoiceC1
	MobilityfromeutracommandCriticalextensionsChoiceCriticalextensionsfuture
)

type MobilityfromeutracommandCriticalextensions struct {
	Choice                   uint64
	C1                       *MobilityfromeutracommandCriticalextensionsC1
	Criticalextensionsfuture *MobilityfromeutracommandCriticalextensionsCriticalextensionsfuture
}
