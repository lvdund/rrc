package ies

import "rrc/utils"

// NR-PRS-MeasurementInfo-r16-nr-MeasPRS-RepetitionAndOffset-r16 ::= CHOICE
// Extensible
const (
	NrPrsMeasurementinfoR16NrMeasprsRepetitionandoffsetR16ChoiceNothing = iota
	NrPrsMeasurementinfoR16NrMeasprsRepetitionandoffsetR16ChoiceMs20R16
	NrPrsMeasurementinfoR16NrMeasprsRepetitionandoffsetR16ChoiceMs40R16
	NrPrsMeasurementinfoR16NrMeasprsRepetitionandoffsetR16ChoiceMs80R16
	NrPrsMeasurementinfoR16NrMeasprsRepetitionandoffsetR16ChoiceMs160R16
)

type NrPrsMeasurementinfoR16NrMeasprsRepetitionandoffsetR16 struct {
	Choice   uint64
	Ms20R16  *utils.INTEGER `lb:0,ub:19`
	Ms40R16  *utils.INTEGER `lb:0,ub:39`
	Ms80R16  *utils.INTEGER `lb:0,ub:79`
	Ms160R16 *utils.INTEGER `lb:0,ub:159`
}
