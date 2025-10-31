package ies

import "rrc/utils"

// BandNR-channelBW-UL-IAB-r16-fr1-100mhz-scs-30kHz ::= ENUMERATED
type BandnrChannelbwUlIabR16Fr1100mhzScs30khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwUlIabR16Fr1100mhzScs30khzEnumeratedNothing = iota
	BandnrChannelbwUlIabR16Fr1100mhzScs30khzEnumeratedSupported
)
