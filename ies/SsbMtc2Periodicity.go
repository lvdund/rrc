package ies

import "rrc/utils"

// SSB-MTC2-periodicity ::= ENUMERATED
type SsbMtc2Periodicity struct {
	Value utils.ENUMERATED
}

const (
	SsbMtc2PeriodicityEnumeratedNothing = iota
	SsbMtc2PeriodicityEnumeratedSf5
	SsbMtc2PeriodicityEnumeratedSf10
	SsbMtc2PeriodicityEnumeratedSf20
	SsbMtc2PeriodicityEnumeratedSf40
	SsbMtc2PeriodicityEnumeratedSf80
	SsbMtc2PeriodicityEnumeratedSpare3
	SsbMtc2PeriodicityEnumeratedSpare2
	SsbMtc2PeriodicityEnumeratedSpare1
)
