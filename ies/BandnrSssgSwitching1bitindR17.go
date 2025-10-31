package ies

import "rrc/utils"

// BandNR-sssg-Switching-1BitInd-r17 ::= ENUMERATED
type BandnrSssgSwitching1bitindR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrSssgSwitching1bitindR17EnumeratedNothing = iota
	BandnrSssgSwitching1bitindR17EnumeratedSupported
)
