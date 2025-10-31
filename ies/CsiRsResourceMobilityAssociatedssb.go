package ies

import "rrc/utils"

// CSI-RS-Resource-Mobility-associatedSSB ::= SEQUENCE
type CsiRsResourceMobilityAssociatedssb struct {
	SsbIndex         SsbIndex
	Isquasicolocated utils.BOOLEAN
}
