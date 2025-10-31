package ies

import "rrc/utils"

// BandNR-channelBW-DL-IAB-r16-fr2-200mhz-scs-120kHz ::= ENUMERATED
type BandnrChannelbwDlIabR16Fr2200mhzScs120khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwDlIabR16Fr2200mhzScs120khzEnumeratedNothing = iota
	BandnrChannelbwDlIabR16Fr2200mhzScs120khzEnumeratedSupported
)
