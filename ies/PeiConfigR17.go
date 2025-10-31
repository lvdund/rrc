package ies

import "rrc/utils"

// PEI-Config-r17 ::= SEQUENCE
// Extensible
type PeiConfigR17 struct {
	PoNumperpeiR17      PeiConfigR17PoNumperpeiR17
	Payloadsizedci27R17 utils.INTEGER `lb:0,ub:maxDCI27SizeR17`
	PeiFrameoffsetR17   utils.INTEGER `lb:0,ub:16`
	SubgroupconfigR17   SubgroupconfigR17
	LastusedcellonlyR17 *utils.BOOLEAN
}
