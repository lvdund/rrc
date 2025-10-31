package ies

// MCCH-MessageType-r17-c1 ::= CHOICE
const (
	McchMessagetypeR17C1ChoiceNothing = iota
	McchMessagetypeR17C1ChoiceMbsbroadcastconfigurationR17
	McchMessagetypeR17C1ChoiceSpare1
)

type McchMessagetypeR17C1 struct {
	Choice                       uint64
	MbsbroadcastconfigurationR17 *MbsbroadcastconfigurationR17
	Spare1                       *struct{}
}
