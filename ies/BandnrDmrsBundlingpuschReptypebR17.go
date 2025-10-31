package ies

import "rrc/utils"

// BandNR-dmrs-BundlingPUSCH-RepTypeB-r17 ::= ENUMERATED
type BandnrDmrsBundlingpuschReptypebR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrDmrsBundlingpuschReptypebR17EnumeratedNothing = iota
	BandnrDmrsBundlingpuschReptypebR17EnumeratedSupported
)
