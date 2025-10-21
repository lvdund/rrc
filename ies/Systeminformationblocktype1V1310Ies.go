package ies

import "rrc/utils"

// SystemInformationBlockType1-v1310-IEs ::= SEQUENCE
type Systeminformationblocktype1V1310Ies struct {
	HypersfnR13                          *utils.BITSTRING
	EdrxAllowedR13                       *utils.ENUMERATED
	CellselectioninfoceR13               *CellselectioninfoceR13
	BandwidthreducedaccessrelatedinfoR13 *interface{}
	Noncriticalextension                 *Systeminformationblocktype1V1320Ies
}
