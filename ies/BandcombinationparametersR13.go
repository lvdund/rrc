package ies

import "rrc/utils"

// BandCombinationParameters-r13 ::= SEQUENCE
type BandcombinationparametersR13 struct {
	DifferentfallbacksupportedR13       *utils.ENUMERATED
	BandparameterlistR13                interface{}
	SupportedbandwidthcombinationsetR13 *SupportedbandwidthcombinationsetR10
	MultipletimingadvanceR13            *utils.ENUMERATED
	SimultaneousrxTxR13                 *utils.ENUMERATED
	BandinfoeutraR13                    Bandinfoeutra
	DcSupportR13                        *interface{}
	Supportednaics2crsApR13             *utils.BITSTRING
	CommsupportedbandsperbcR13          *utils.BITSTRING
}
