package ies

// BandNR-channelBWs-UL ::= CHOICE
const (
	BandnrChannelbwsUlChoiceNothing = iota
	BandnrChannelbwsUlChoiceFr1
	BandnrChannelbwsUlChoiceFr2
)

type BandnrChannelbwsUl struct {
	Choice uint64
	Fr1    *BandnrChannelbwsUlFr1
	Fr2    *BandnrChannelbwsUlFr2
}
