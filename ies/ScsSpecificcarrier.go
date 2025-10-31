package ies

import "rrc/utils"

// SCS-SpecificCarrier ::= SEQUENCE
// Extensible
type ScsSpecificcarrier struct {
	Offsettocarrier         utils.INTEGER `lb:0,ub:2199`
	Subcarrierspacing       Subcarrierspacing
	Carrierbandwidth        utils.INTEGER  `lb:0,ub:maxNrofPhysicalResourceBlocks`
	Txdirectcurrentlocation *utils.INTEGER `lb:0,ub:4095,ext`
}
