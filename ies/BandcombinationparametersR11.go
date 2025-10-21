package ies

import "rrc/utils"

// BandCombinationParameters-r11 ::= SEQUENCE
// Extensible
type BandcombinationparametersR11 struct {
	BandparameterlistR11                interface{}
	SupportedbandwidthcombinationsetR11 *SupportedbandwidthcombinationsetR10
	MultipletimingadvanceR11            *utils.ENUMERATED
	SimultaneousrxTxR11                 *utils.ENUMERATED
	BandinfoeutraR11                    Bandinfoeutra
}
