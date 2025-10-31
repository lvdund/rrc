package ies

import "rrc/utils"

// SRS-ConfigAp-r13 ::= SEQUENCE
type SrsConfigapR13 struct {
	SrsAntennaportapR13     SrsAntennaport
	SrsBandwidthapR13       SrsConfigapR13SrsBandwidthapR13
	FreqdomainpositionapR13 utils.INTEGER `lb:0,ub:23`
	TransmissioncombapR13   utils.INTEGER `lb:0,ub:3`
	CyclicshiftapR13        SrsConfigapR13CyclicshiftapR13
	TransmissioncombnumR13  SrsConfigapR13TransmissioncombnumR13
}
