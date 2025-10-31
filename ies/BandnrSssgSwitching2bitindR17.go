package ies

import "rrc/utils"

// BandNR-sssg-Switching-2BitInd-r17 ::= ENUMERATED
type BandnrSssgSwitching2bitindR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrSssgSwitching2bitindR17EnumeratedNothing = iota
	BandnrSssgSwitching2bitindR17EnumeratedSupported
)
