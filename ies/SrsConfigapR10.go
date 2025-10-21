package ies

import "rrc/utils"

// SRS-ConfigAp-r10 ::= SEQUENCE
type SrsConfigapR10 struct {
	SrsAntennaportapR10     SrsAntennaport
	SrsBandwidthapR10       utils.ENUMERATED
	FreqdomainpositionapR10 utils.INTEGER
	TransmissioncombapR10   utils.INTEGER
	CyclicshiftapR10        utils.ENUMERATED
}
