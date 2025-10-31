package ies

import "rrc/utils"

// BandNR-channelBW-UL-IAB-r16-fr2-200mhz-scs-60kHz ::= ENUMERATED
type BandnrChannelbwUlIabR16Fr2200mhzScs60khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwUlIabR16Fr2200mhzScs60khzEnumeratedNothing = iota
	BandnrChannelbwUlIabR16Fr2200mhzScs60khzEnumeratedSupported
)
