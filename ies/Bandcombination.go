package ies

import "rrc/utils"

// BandCombination ::= SEQUENCE
type Bandcombination struct {
	Bandlist                         []Bandparameters `lb:1,ub:maxSimultaneousBands`
	Featuresetcombination            Featuresetcombinationid
	CaParameterseutra                *CaParameterseutra
	CaParametersnr                   *CaParametersnr
	MrdcParameters                   *MrdcParameters
	Supportedbandwidthcombinationset *utils.BITSTRING `lb:1,ub:32`
	PowerclassV1530                  *BandcombinationPowerclassV1530
}
