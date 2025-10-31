package ies

import "rrc/utils"

// BandNR-interSlotFreqHopPUCCH-r17 ::= ENUMERATED
type BandnrInterslotfreqhoppucchR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrInterslotfreqhoppucchR17EnumeratedNothing = iota
	BandnrInterslotfreqhoppucchR17EnumeratedSupported
)
