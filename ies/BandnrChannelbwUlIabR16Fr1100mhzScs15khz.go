package ies

import "rrc/utils"

// BandNR-channelBW-UL-IAB-r16-fr1-100mhz-scs-15kHz ::= ENUMERATED
type BandnrChannelbwUlIabR16Fr1100mhzScs15khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwUlIabR16Fr1100mhzScs15khzEnumeratedNothing = iota
	BandnrChannelbwUlIabR16Fr1100mhzScs15khzEnumeratedSupported
)
