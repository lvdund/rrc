package ies

import "rrc/utils"

// PCCH-Config-nB ::= ENUMERATED
type PcchConfigNb struct {
	Value utils.ENUMERATED
}

const (
	PcchConfigNbEnumeratedNothing = iota
	PcchConfigNbEnumeratedFourt
	PcchConfigNbEnumeratedTwot
	PcchConfigNbEnumeratedOnet
	PcchConfigNbEnumeratedHalft
	PcchConfigNbEnumeratedQuartert
	PcchConfigNbEnumeratedOneeightht
	PcchConfigNbEnumeratedOnesixteentht
	PcchConfigNbEnumeratedOnethirtysecondt
)
