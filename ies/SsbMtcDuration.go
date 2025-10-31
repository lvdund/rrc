package ies

import "rrc/utils"

// SSB-MTC-duration ::= ENUMERATED
type SsbMtcDuration struct {
	Value utils.ENUMERATED
}

const (
	SsbMtcDurationEnumeratedNothing = iota
	SsbMtcDurationEnumeratedSf1
	SsbMtcDurationEnumeratedSf2
	SsbMtcDurationEnumeratedSf3
	SsbMtcDurationEnumeratedSf4
	SsbMtcDurationEnumeratedSf5
)
