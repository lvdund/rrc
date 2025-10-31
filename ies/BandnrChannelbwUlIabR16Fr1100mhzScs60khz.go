package ies

import "rrc/utils"

// BandNR-channelBW-UL-IAB-r16-fr1-100mhz-scs-60kHz ::= ENUMERATED
type BandnrChannelbwUlIabR16Fr1100mhzScs60khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwUlIabR16Fr1100mhzScs60khzEnumeratedNothing = iota
	BandnrChannelbwUlIabR16Fr1100mhzScs60khzEnumeratedSupported
)
