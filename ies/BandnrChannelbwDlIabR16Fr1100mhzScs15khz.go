package ies

import "rrc/utils"

// BandNR-channelBW-DL-IAB-r16-fr1-100mhz-scs-15kHz ::= ENUMERATED
type BandnrChannelbwDlIabR16Fr1100mhzScs15khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwDlIabR16Fr1100mhzScs15khzEnumeratedNothing = iota
	BandnrChannelbwDlIabR16Fr1100mhzScs15khzEnumeratedSupported
)
