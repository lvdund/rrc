package ies

// BandNR-channelBW-DL-IAB-r16 ::= CHOICE
const (
	BandnrChannelbwDlIabR16ChoiceNothing = iota
	BandnrChannelbwDlIabR16ChoiceFr1100mhz
	BandnrChannelbwDlIabR16ChoiceFr2200mhz
)

type BandnrChannelbwDlIabR16 struct {
	Choice    uint64
	Fr1100mhz *BandnrChannelbwDlIabR16Fr1100mhz
	Fr2200mhz *BandnrChannelbwDlIabR16Fr2200mhz
}
