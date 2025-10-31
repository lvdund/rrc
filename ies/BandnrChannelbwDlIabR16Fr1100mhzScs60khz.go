package ies

import "rrc/utils"

// BandNR-channelBW-DL-IAB-r16-fr1-100mhz-scs-60kHz ::= ENUMERATED
type BandnrChannelbwDlIabR16Fr1100mhzScs60khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwDlIabR16Fr1100mhzScs60khzEnumeratedNothing = iota
	BandnrChannelbwDlIabR16Fr1100mhzScs60khzEnumeratedSupported
)
