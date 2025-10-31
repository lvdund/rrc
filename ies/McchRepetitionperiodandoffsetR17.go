package ies

import "rrc/utils"

// MCCH-RepetitionPeriodAndOffset-r17 ::= CHOICE
const (
	McchRepetitionperiodandoffsetR17ChoiceNothing = iota
	McchRepetitionperiodandoffsetR17ChoiceRf1R17
	McchRepetitionperiodandoffsetR17ChoiceRf2R17
	McchRepetitionperiodandoffsetR17ChoiceRf4R17
	McchRepetitionperiodandoffsetR17ChoiceRf8R17
	McchRepetitionperiodandoffsetR17ChoiceRf16R17
	McchRepetitionperiodandoffsetR17ChoiceRf32R17
	McchRepetitionperiodandoffsetR17ChoiceRf64R17
	McchRepetitionperiodandoffsetR17ChoiceRf128R17
	McchRepetitionperiodandoffsetR17ChoiceRf256R17
)

type McchRepetitionperiodandoffsetR17 struct {
	Choice   uint64
	Rf1R17   *utils.INTEGER `lb:0,ub:0`
	Rf2R17   *utils.INTEGER `lb:0,ub:1`
	Rf4R17   *utils.INTEGER `lb:0,ub:3`
	Rf8R17   *utils.INTEGER `lb:0,ub:7`
	Rf16R17  *utils.INTEGER `lb:0,ub:15`
	Rf32R17  *utils.INTEGER `lb:0,ub:31`
	Rf64R17  *utils.INTEGER `lb:0,ub:63`
	Rf128R17 *utils.INTEGER `lb:0,ub:127`
	Rf256R17 *utils.INTEGER `lb:0,ub:255`
}
