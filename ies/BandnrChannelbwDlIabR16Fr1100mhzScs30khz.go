package ies

import "rrc/utils"

// BandNR-channelBW-DL-IAB-r16-fr1-100mhz-scs-30kHz ::= ENUMERATED
type BandnrChannelbwDlIabR16Fr1100mhzScs30khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwDlIabR16Fr1100mhzScs30khzEnumeratedNothing = iota
	BandnrChannelbwDlIabR16Fr1100mhzScs30khzEnumeratedSupported
)
