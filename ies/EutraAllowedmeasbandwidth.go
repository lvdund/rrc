package ies

import "rrc/utils"

// EUTRA-AllowedMeasBandwidth ::= ENUMERATED
type EutraAllowedmeasbandwidth struct {
	Value utils.ENUMERATED
}

const (
	EutraAllowedmeasbandwidthEnumeratedNothing = iota
	EutraAllowedmeasbandwidthEnumeratedMbw6
	EutraAllowedmeasbandwidthEnumeratedMbw15
	EutraAllowedmeasbandwidthEnumeratedMbw25
	EutraAllowedmeasbandwidthEnumeratedMbw50
	EutraAllowedmeasbandwidthEnumeratedMbw75
	EutraAllowedmeasbandwidthEnumeratedMbw100
)
