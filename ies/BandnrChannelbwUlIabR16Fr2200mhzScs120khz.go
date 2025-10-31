package ies

import "rrc/utils"

// BandNR-channelBW-UL-IAB-r16-fr2-200mhz-scs-120kHz ::= ENUMERATED
type BandnrChannelbwUlIabR16Fr2200mhzScs120khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwUlIabR16Fr2200mhzScs120khzEnumeratedNothing = iota
	BandnrChannelbwUlIabR16Fr2200mhzScs120khzEnumeratedSupported
)
