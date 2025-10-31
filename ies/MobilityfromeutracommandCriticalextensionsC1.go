package ies

// MobilityFromEUTRACommand-criticalExtensions-c1 ::= CHOICE
const (
	MobilityfromeutracommandCriticalextensionsC1ChoiceNothing = iota
	MobilityfromeutracommandCriticalextensionsC1ChoiceMobilityfromeutracommandR8
	MobilityfromeutracommandCriticalextensionsC1ChoiceMobilityfromeutracommandR9
	MobilityfromeutracommandCriticalextensionsC1ChoiceSpare2
	MobilityfromeutracommandCriticalextensionsC1ChoiceSpare1
)

type MobilityfromeutracommandCriticalextensionsC1 struct {
	Choice                     uint64
	MobilityfromeutracommandR8 *MobilityfromeutracommandR8
	MobilityfromeutracommandR9 *MobilityfromeutracommandR9
	Spare2                     *struct{}
	Spare1                     *struct{}
}
