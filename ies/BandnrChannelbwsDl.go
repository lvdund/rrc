package ies

// BandNR-channelBWs-DL ::= CHOICE
const (
	BandnrChannelbwsDlChoiceNothing = iota
	BandnrChannelbwsDlChoiceFr1
	BandnrChannelbwsDlChoiceFr2
)

type BandnrChannelbwsDl struct {
	Choice uint64
	Fr1    *BandnrChannelbwsDlFr1
	Fr2    *BandnrChannelbwsDlFr2
}
