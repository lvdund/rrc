package ies

import "rrc/utils"

// SRS-ConfigAp-r10-srs-BandwidthAp-r10 ::= ENUMERATED
type SrsConfigapR10SrsBandwidthapR10 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigapR10SrsBandwidthapR10EnumeratedNothing = iota
	SrsConfigapR10SrsBandwidthapR10EnumeratedBw0
	SrsConfigapR10SrsBandwidthapR10EnumeratedBw1
	SrsConfigapR10SrsBandwidthapR10EnumeratedBw2
	SrsConfigapR10SrsBandwidthapR10EnumeratedBw3
)
