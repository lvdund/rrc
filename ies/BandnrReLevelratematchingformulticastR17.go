package ies

import "rrc/utils"

// BandNR-re-LevelRateMatchingForMulticast-r17 ::= ENUMERATED
type BandnrReLevelratematchingformulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrReLevelratematchingformulticastR17EnumeratedNothing = iota
	BandnrReLevelratematchingformulticastR17EnumeratedSupported
)
