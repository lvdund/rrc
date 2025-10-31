package ies

import "rrc/utils"

// CA-ParametersEUTRA ::= SEQUENCE
// Extensible
type CaParameterseutra struct {
	Multipletimingadvance                      *CaParameterseutraMultipletimingadvance
	SimultaneousrxTx                           *CaParameterseutraSimultaneousrxTx
	Supportednaics2crsAp                       *utils.BITSTRING `lb:1,ub:8`
	AdditionalrxTxPerformancereq               *CaParameterseutraAdditionalrxTxPerformancereq
	UeCaPowerclassN                            *CaParameterseutraUeCaPowerclassN
	SupportedbandwidthcombinationseteutraV1530 *utils.BITSTRING `lb:1,ub:32`
}
