package ies

import "rrc/utils"

// SRS-ConfigAp-r13-srs-BandwidthAp-r13 ::= ENUMERATED
type SrsConfigapR13SrsBandwidthapR13 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigapR13SrsBandwidthapR13EnumeratedNothing = iota
	SrsConfigapR13SrsBandwidthapR13EnumeratedBw0
	SrsConfigapR13SrsBandwidthapR13EnumeratedBw1
	SrsConfigapR13SrsBandwidthapR13EnumeratedBw2
	SrsConfigapR13SrsBandwidthapR13EnumeratedBw3
)
