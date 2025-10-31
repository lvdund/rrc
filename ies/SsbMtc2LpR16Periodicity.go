package ies

import "rrc/utils"

// SSB-MTC2-LP-r16-periodicity ::= ENUMERATED
type SsbMtc2LpR16Periodicity struct {
	Value utils.ENUMERATED
}

const (
	SsbMtc2LpR16PeriodicityEnumeratedNothing = iota
	SsbMtc2LpR16PeriodicityEnumeratedSf10
	SsbMtc2LpR16PeriodicityEnumeratedSf20
	SsbMtc2LpR16PeriodicityEnumeratedSf40
	SsbMtc2LpR16PeriodicityEnumeratedSf80
	SsbMtc2LpR16PeriodicityEnumeratedSf160
	SsbMtc2LpR16PeriodicityEnumeratedSpare3
	SsbMtc2LpR16PeriodicityEnumeratedSpare2
	SsbMtc2LpR16PeriodicityEnumeratedSpare1
)
