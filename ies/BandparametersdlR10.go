package ies

import "rrc/utils"

// BandParametersDL-r10 ::= SEQUENCE OF CA-MIMO-ParametersDL-r10
// SIZE (1..maxBandwidthClass-r10)
type BandparametersdlR10 struct {
	Value []CaMimoParametersdlR10
}
