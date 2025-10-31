package ies

import "rrc/utils"

// BandNR-ue-specific-K-Offset-r17 ::= ENUMERATED
type BandnrUeSpecificKOffsetR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrUeSpecificKOffsetR17EnumeratedNothing = iota
	BandnrUeSpecificKOffsetR17EnumeratedSupported
)
