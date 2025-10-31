package ies

import "rrc/utils"

// BandNR-eventA4BasedCondHandover-r17 ::= ENUMERATED
type BandnrEventa4basedcondhandoverR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrEventa4basedcondhandoverR17EnumeratedNothing = iota
	BandnrEventa4basedcondhandoverR17EnumeratedSupported
)
