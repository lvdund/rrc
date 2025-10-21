package ies

import "rrc/utils"

// SupportedBandInfoList-r12 ::= SEQUENCE OF SupportedBandInfo-r12
// SIZE (1..maxBands)
type SupportedbandinfolistR12 struct {
	Value []SupportedbandinfoR12
}
