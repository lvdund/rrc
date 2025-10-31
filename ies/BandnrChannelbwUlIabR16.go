package ies

// BandNR-channelBW-UL-IAB-r16 ::= CHOICE
const (
	BandnrChannelbwUlIabR16ChoiceNothing = iota
	BandnrChannelbwUlIabR16ChoiceFr1100mhz
	BandnrChannelbwUlIabR16ChoiceFr2200mhz
)

type BandnrChannelbwUlIabR16 struct {
	Choice    uint64
	Fr1100mhz *BandnrChannelbwUlIabR16Fr1100mhz
	Fr2200mhz *BandnrChannelbwUlIabR16Fr2200mhz
}
