package ies

import "rrc/utils"

// DiscardTimerExt-r16 ::= ENUMERATED
type DiscardtimerextR16 struct {
	Value utils.ENUMERATED
}

const (
	DiscardtimerextR16EnumeratedNothing = iota
	DiscardtimerextR16EnumeratedMs0dot5
	DiscardtimerextR16EnumeratedMs1
	DiscardtimerextR16EnumeratedMs2
	DiscardtimerextR16EnumeratedMs4
	DiscardtimerextR16EnumeratedMs6
	DiscardtimerextR16EnumeratedMs8
	DiscardtimerextR16EnumeratedSpare2
	DiscardtimerextR16EnumeratedSpare1
)
