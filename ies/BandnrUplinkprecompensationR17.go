package ies

import "rrc/utils"

// BandNR-uplinkPreCompensation-r17 ::= ENUMERATED
type BandnrUplinkprecompensationR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrUplinkprecompensationR17EnumeratedNothing = iota
	BandnrUplinkprecompensationR17EnumeratedSupported
)
