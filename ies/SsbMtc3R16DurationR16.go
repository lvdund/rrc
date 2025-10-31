package ies

import "rrc/utils"

// SSB-MTC3-r16-duration-r16 ::= ENUMERATED
type SsbMtc3R16DurationR16 struct {
	Value utils.ENUMERATED
}

const (
	SsbMtc3R16DurationR16EnumeratedNothing = iota
	SsbMtc3R16DurationR16EnumeratedSf1
	SsbMtc3R16DurationR16EnumeratedSf2
	SsbMtc3R16DurationR16EnumeratedSf3
	SsbMtc3R16DurationR16EnumeratedSf4
	SsbMtc3R16DurationR16EnumeratedSf5
)
