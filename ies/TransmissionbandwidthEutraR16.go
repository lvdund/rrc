package ies

import "rrc/utils"

// TransmissionBandwidth-EUTRA-r16 ::= ENUMERATED
type TransmissionbandwidthEutraR16 struct {
	Value utils.ENUMERATED
}

const (
	TransmissionbandwidthEutraR16EnumeratedNothing = iota
	TransmissionbandwidthEutraR16EnumeratedRb6
	TransmissionbandwidthEutraR16EnumeratedRb15
	TransmissionbandwidthEutraR16EnumeratedRb25
	TransmissionbandwidthEutraR16EnumeratedRb50
	TransmissionbandwidthEutraR16EnumeratedRb75
	TransmissionbandwidthEutraR16EnumeratedRb100
)
