package ies

import "rrc/utils"

// AllowedMeasBandwidth ::= ENUMERATED
type Allowedmeasbandwidth struct {
	Value utils.ENUMERATED
}

const (
	AllowedmeasbandwidthEnumeratedNothing = iota
	AllowedmeasbandwidthEnumeratedMbw6
	AllowedmeasbandwidthEnumeratedMbw15
	AllowedmeasbandwidthEnumeratedMbw25
	AllowedmeasbandwidthEnumeratedMbw50
	AllowedmeasbandwidthEnumeratedMbw75
	AllowedmeasbandwidthEnumeratedMbw100
)
