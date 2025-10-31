package ies

import "rrc/utils"

// BandNR-enhancedSkipUplinkTxDynamic-v1660 ::= ENUMERATED
type BandnrEnhancedskipuplinktxdynamicV1660 struct {
	Value utils.ENUMERATED
}

const (
	BandnrEnhancedskipuplinktxdynamicV1660EnumeratedNothing = iota
	BandnrEnhancedskipuplinktxdynamicV1660EnumeratedSupported
)
