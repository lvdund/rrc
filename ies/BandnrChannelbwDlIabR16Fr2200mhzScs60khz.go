package ies

import "rrc/utils"

// BandNR-channelBW-DL-IAB-r16-fr2-200mhz-scs-60kHz ::= ENUMERATED
type BandnrChannelbwDlIabR16Fr2200mhzScs60khz struct {
	Value utils.ENUMERATED
}

const (
	BandnrChannelbwDlIabR16Fr2200mhzScs60khzEnumeratedNothing = iota
	BandnrChannelbwDlIabR16Fr2200mhzScs60khzEnumeratedSupported
)
