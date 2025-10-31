package ies

import "rrc/utils"

// GNSS-ID-r15-gnss-id-r15 ::= utils.ENUMERATED // Extensible
type GnssIdR15GnssIdR15 struct {
	Value utils.ENUMERATED
}

const (
	GnssIdR15GnssIdR15EnumeratedNothing = iota
	GnssIdR15GnssIdR15EnumeratedGps
	GnssIdR15GnssIdR15EnumeratedSbas
	GnssIdR15GnssIdR15EnumeratedQzss
	GnssIdR15GnssIdR15EnumeratedGalileo
	GnssIdR15GnssIdR15EnumeratedGlonass
	GnssIdR15GnssIdR15EnumeratedBds
	GnssIdR15GnssIdR15EnumeratedNavic_V1610
)
