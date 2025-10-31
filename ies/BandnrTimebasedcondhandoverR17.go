package ies

import "rrc/utils"

// BandNR-timeBasedCondHandover-r17 ::= ENUMERATED
type BandnrTimebasedcondhandoverR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrTimebasedcondhandoverR17EnumeratedNothing = iota
	BandnrTimebasedcondhandoverR17EnumeratedSupported
)
