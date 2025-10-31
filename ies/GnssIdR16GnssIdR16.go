package ies

import "rrc/utils"

// GNSS-ID-r16-gnss-id-r16 ::= utils.ENUMERATED // Extensible
type GnssIdR16GnssIdR16 struct {
	Value utils.ENUMERATED
}

const (
	GnssIdR16GnssIdR16EnumeratedNothing = iota
	GnssIdR16GnssIdR16EnumeratedGps
	GnssIdR16GnssIdR16EnumeratedSbas
	GnssIdR16GnssIdR16EnumeratedQzss
	GnssIdR16GnssIdR16EnumeratedGalileo
	GnssIdR16GnssIdR16EnumeratedGlonass
	GnssIdR16GnssIdR16EnumeratedBds
)
