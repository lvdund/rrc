package ies

import "rrc/utils"

// SupportedBandEUTRA ::= SEQUENCE
type Supportedbandeutra struct {
	Bandeutra  Freqbandindicator
	Halfduplex utils.BOOLEAN
}
