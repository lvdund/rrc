package ies

// MCCH-MessageType-later-c2 ::= CHOICE
const (
	McchMessagetypeLaterC2ChoiceNothing = iota
	McchMessagetypeLaterC2ChoiceMbmscountingrequestR10
)

type McchMessagetypeLaterC2 struct {
	Choice                 uint64
	MbmscountingrequestR10 *MbmscountingrequestR10
}
