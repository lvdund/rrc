package ies

import "rrc/utils"

// SRS-ConfigAp-r10 ::= SEQUENCE
type SrsConfigapR10 struct {
	SrsAntennaportapR10     SrsAntennaport
	SrsBandwidthapR10       SrsConfigapR10SrsBandwidthapR10
	FreqdomainpositionapR10 utils.INTEGER `lb:0,ub:23`
	TransmissioncombapR10   utils.INTEGER `lb:0,ub:1`
	CyclicshiftapR10        SrsConfigapR10CyclicshiftapR10
}
