package ies

import "rrc/utils"

// MTC-SSB-NR-r15 ::= SEQUENCE
type MtcSsbNrR15 struct {
	PeriodicityandoffsetR15 interface{}
	SsbDurationR15          utils.ENUMERATED
}
