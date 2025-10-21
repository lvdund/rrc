package ies

import "rrc/utils"

// SRS-ConfigAp-r13 ::= SEQUENCE
type SrsConfigapR13 struct {
	SrsAntennaportapR13     SrsAntennaport
	SrsBandwidthapR13       utils.ENUMERATED
	FreqdomainpositionapR13 utils.INTEGER
	TransmissioncombapR13   utils.INTEGER
	CyclicshiftapR13        utils.ENUMERATED
	TransmissioncombnumR13  utils.ENUMERATED
}
