package ies

import "rrc/utils"

// BandNR-dynamicMulticastDCI-Format4-2-r17 ::= ENUMERATED
type BandnrDynamicmulticastdciFormat42R17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrDynamicmulticastdciFormat42R17EnumeratedNothing = iota
	BandnrDynamicmulticastdciFormat42R17EnumeratedSupported
)
