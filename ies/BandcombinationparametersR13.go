package ies

import "rrc/utils"

// BandCombinationParameters-r13 ::= SEQUENCE
type BandcombinationparametersR13 struct {
	DifferentfallbacksupportedR13       *bool
	BandparameterlistR13                []BandparametersR13 `lb:1,ub:maxSimultaneousBandsR10`
	SupportedbandwidthcombinationsetR13 *SupportedbandwidthcombinationsetR10
	MultipletimingadvanceR13            *BandcombinationparametersR13MultipletimingadvanceR13
	SimultaneousrxTxR13                 *BandcombinationparametersR13SimultaneousrxTxR13
	BandinfoeutraR13                    Bandinfoeutra
	DcSupportR13                        *BandcombinationparametersR13DcSupportR13
	Supportednaics2crsApR13             *utils.BITSTRING `lb:1,ub:maxNAICSEntriesR12`
	CommsupportedbandsperbcR13          *utils.BITSTRING `lb:1,ub:maxBands`
}
