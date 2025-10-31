package ies

import "rrc/utils"

// BandNR-dmrs-BundlingPUSCH-RepTypeA-r17 ::= ENUMERATED
type BandnrDmrsBundlingpuschReptypeaR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrDmrsBundlingpuschReptypeaR17EnumeratedNothing = iota
	BandnrDmrsBundlingpuschReptypeaR17EnumeratedSupported
)
