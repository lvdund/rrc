package ies

import "rrc/utils"

// BandNR-locationBasedCondHandover-r17 ::= ENUMERATED
type BandnrLocationbasedcondhandoverR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrLocationbasedcondhandoverR17EnumeratedNothing = iota
	BandnrLocationbasedcondhandoverR17EnumeratedSupported
)
