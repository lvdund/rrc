package ies

// BandNR-channelBWs-DL-v1590 ::= CHOICE
const (
	BandnrChannelbwsDlV1590ChoiceNothing = iota
	BandnrChannelbwsDlV1590ChoiceFr1
	BandnrChannelbwsDlV1590ChoiceFr2
)

type BandnrChannelbwsDlV1590 struct {
	Choice uint64
	Fr1    *BandnrChannelbwsDlV1590Fr1
	Fr2    *BandnrChannelbwsDlV1590Fr2
}
