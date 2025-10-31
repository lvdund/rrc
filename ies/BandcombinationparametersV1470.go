package ies

import "rrc/utils"

// BandCombinationParameters-v1470 ::= SEQUENCE
type BandcombinationparametersV1470 struct {
	BandparameterlistV1470   *[]BandparametersV1470 `lb:1,ub:maxSimultaneousBandsR10`
	SrsMaxsimultaneousccsR14 *utils.INTEGER         `lb:0,ub:31`
}
