package ies

import "rrc/utils"

// DiscardTimerExt2-r17 ::= ENUMERATED
type Discardtimerext2R17 struct {
	Value utils.ENUMERATED
}

const (
	Discardtimerext2R17EnumeratedNothing = iota
	Discardtimerext2R17EnumeratedMs2000
	Discardtimerext2R17EnumeratedSpare3
	Discardtimerext2R17EnumeratedSpare2
	Discardtimerext2R17EnumeratedSpare1
)
