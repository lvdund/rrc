package ies

// FreqBandInformation ::= CHOICE
const (
	FreqbandinformationChoiceNothing = iota
	FreqbandinformationChoiceBandinformationeutra
	FreqbandinformationChoiceBandinformationnr
)

type Freqbandinformation struct {
	Choice               uint64
	Bandinformationeutra *Freqbandinformationeutra
	Bandinformationnr    *Freqbandinformationnr
}
