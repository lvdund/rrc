package ies

import "rrc/utils"

// NR-DL-PRS-Periodicity-and-ResourceSetSlotOffset-r17-scs15-r17 ::= CHOICE
// Extensible
const (
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceNothing = iota
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN4R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN5R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN8R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN10R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN16R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN20R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN32R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN40R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN64R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN80R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN160R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN320R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN640R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN1280R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN2560R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN5120R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17ChoiceN10240R17
)

type NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs15R17 struct {
	Choice    uint64
	N4R17     *utils.INTEGER `lb:0,ub:3`
	N5R17     *utils.INTEGER `lb:0,ub:4`
	N8R17     *utils.INTEGER `lb:0,ub:7`
	N10R17    *utils.INTEGER `lb:0,ub:9`
	N16R17    *utils.INTEGER `lb:0,ub:15`
	N20R17    *utils.INTEGER `lb:0,ub:19`
	N32R17    *utils.INTEGER `lb:0,ub:31`
	N40R17    *utils.INTEGER `lb:0,ub:39`
	N64R17    *utils.INTEGER `lb:0,ub:63`
	N80R17    *utils.INTEGER `lb:0,ub:79`
	N160R17   *utils.INTEGER `lb:0,ub:159`
	N320R17   *utils.INTEGER `lb:0,ub:319`
	N640R17   *utils.INTEGER `lb:0,ub:639`
	N1280R17  *utils.INTEGER `lb:0,ub:1279`
	N2560R17  *utils.INTEGER `lb:0,ub:2559`
	N5120R17  *utils.INTEGER `lb:0,ub:5119`
	N10240R17 *utils.INTEGER `lb:0,ub:10239`
}
