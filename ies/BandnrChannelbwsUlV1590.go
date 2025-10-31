package ies

// BandNR-channelBWs-UL-v1590 ::= CHOICE
const (
	BandnrChannelbwsUlV1590ChoiceNothing = iota
	BandnrChannelbwsUlV1590ChoiceFr1
	BandnrChannelbwsUlV1590ChoiceFr2
)

type BandnrChannelbwsUlV1590 struct {
	Choice uint64
	Fr1    *BandnrChannelbwsUlV1590Fr1
	Fr2    *BandnrChannelbwsUlV1590Fr2
}
