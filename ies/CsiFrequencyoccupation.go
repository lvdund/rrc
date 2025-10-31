package ies

import "rrc/utils"

// CSI-FrequencyOccupation ::= SEQUENCE
// Extensible
type CsiFrequencyoccupation struct {
	Startingrb utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocks1`
	Nrofrbs    utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocksPlus1`
}
