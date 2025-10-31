package ies

// DL-PPW-PeriodicityAndStartSlot-r17 ::= CHOICE
// Extensible
const (
	DlPpwPeriodicityandstartslotR17ChoiceNothing = iota
	DlPpwPeriodicityandstartslotR17ChoiceScs15
	DlPpwPeriodicityandstartslotR17ChoiceScs30
	DlPpwPeriodicityandstartslotR17ChoiceScs60
	DlPpwPeriodicityandstartslotR17ChoiceScs120
)

type DlPpwPeriodicityandstartslotR17 struct {
	Choice uint64
	Scs15  *DlPpwPeriodicityandstartslotR17Scs15
	Scs30  *DlPpwPeriodicityandstartslotR17Scs30
	Scs60  *DlPpwPeriodicityandstartslotR17Scs60
	Scs120 *DlPpwPeriodicityandstartslotR17Scs120
}
