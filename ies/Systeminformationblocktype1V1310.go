package ies

import "rrc/utils"

// SystemInformationBlockType1-v1310-IEs ::= SEQUENCE
type Systeminformationblocktype1V1310 struct {
	HypersfnR13                          *utils.BITSTRING `lb:10,ub:10`
	EdrxAllowedR13                       *bool
	CellselectioninfoceR13               *CellselectioninfoceR13
	BandwidthreducedaccessrelatedinfoR13 *Systeminformationblocktype1V1310IesBandwidthreducedaccessrelatedinfoR13
	Noncriticalextension                 *Systeminformationblocktype1V1320
}
