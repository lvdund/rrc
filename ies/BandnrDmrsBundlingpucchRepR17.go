package ies

import "rrc/utils"

// BandNR-dmrs-BundlingPUCCH-Rep-r17 ::= ENUMERATED
type BandnrDmrsBundlingpucchRepR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrDmrsBundlingpucchRepR17EnumeratedNothing = iota
	BandnrDmrsBundlingpucchRepR17EnumeratedSupported
)
