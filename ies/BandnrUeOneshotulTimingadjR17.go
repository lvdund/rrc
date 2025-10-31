package ies

import "rrc/utils"

// BandNR-ue-OneShotUL-TimingAdj-r17 ::= ENUMERATED
type BandnrUeOneshotulTimingadjR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrUeOneshotulTimingadjR17EnumeratedNothing = iota
	BandnrUeOneshotulTimingadjR17EnumeratedSupported
)
