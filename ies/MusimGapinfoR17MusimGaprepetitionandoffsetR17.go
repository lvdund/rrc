package ies

import "rrc/utils"

// MUSIM-GapInfo-r17-musim-GapRepetitionAndOffset-r17 ::= CHOICE
// Extensible
const (
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceNothing = iota
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceMs20R17
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceMs40R17
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceMs80R17
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceMs160R17
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceMs320R17
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceMs640R17
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceMs1280R17
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceMs2560R17
	MusimGapinfoR17MusimGaprepetitionandoffsetR17ChoiceMs5120R17
)

type MusimGapinfoR17MusimGaprepetitionandoffsetR17 struct {
	Choice    uint64
	Ms20R17   *utils.INTEGER `lb:0,ub:19`
	Ms40R17   *utils.INTEGER `lb:0,ub:39`
	Ms80R17   *utils.INTEGER `lb:0,ub:79`
	Ms160R17  *utils.INTEGER `lb:0,ub:159`
	Ms320R17  *utils.INTEGER `lb:0,ub:319`
	Ms640R17  *utils.INTEGER `lb:0,ub:639`
	Ms1280R17 *utils.INTEGER `lb:0,ub:1279`
	Ms2560R17 *utils.INTEGER `lb:0,ub:2559`
	Ms5120R17 *utils.INTEGER `lb:0,ub:5119`
}
