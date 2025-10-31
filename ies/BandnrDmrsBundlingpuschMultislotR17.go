package ies

import "rrc/utils"

// BandNR-dmrs-BundlingPUSCH-multiSlot-r17 ::= ENUMERATED
type BandnrDmrsBundlingpuschMultislotR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrDmrsBundlingpuschMultislotR17EnumeratedNothing = iota
	BandnrDmrsBundlingpuschMultislotR17EnumeratedSupported
)
