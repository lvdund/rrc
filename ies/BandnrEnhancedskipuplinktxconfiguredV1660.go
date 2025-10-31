package ies

import "rrc/utils"

// BandNR-enhancedSkipUplinkTxConfigured-v1660 ::= ENUMERATED
type BandnrEnhancedskipuplinktxconfiguredV1660 struct {
	Value utils.ENUMERATED
}

const (
	BandnrEnhancedskipuplinktxconfiguredV1660EnumeratedNothing = iota
	BandnrEnhancedskipuplinktxconfiguredV1660EnumeratedSupported
)
