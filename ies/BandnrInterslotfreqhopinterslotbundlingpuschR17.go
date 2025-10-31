package ies

import "rrc/utils"

// BandNR-interSlotFreqHopInterSlotBundlingPUSCH-r17 ::= ENUMERATED
type BandnrInterslotfreqhopinterslotbundlingpuschR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrInterslotfreqhopinterslotbundlingpuschR17EnumeratedNothing = iota
	BandnrInterslotfreqhopinterslotbundlingpuschR17EnumeratedSupported
)
