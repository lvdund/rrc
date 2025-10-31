package ies

import "rrc/utils"

// BandNR-dmrs-BundlingNonBackToBackTX-r17 ::= ENUMERATED
type BandnrDmrsBundlingnonbacktobacktxR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrDmrsBundlingnonbacktobacktxR17EnumeratedNothing = iota
	BandnrDmrsBundlingnonbacktobacktxR17EnumeratedSupported
)
