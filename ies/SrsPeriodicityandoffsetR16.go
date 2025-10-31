package ies

import "rrc/utils"

// SRS-PeriodicityAndOffset-r16 ::= CHOICE
// Extensible
const (
	SrsPeriodicityandoffsetR16ChoiceNothing = iota
	SrsPeriodicityandoffsetR16ChoiceSl1
	SrsPeriodicityandoffsetR16ChoiceSl2
	SrsPeriodicityandoffsetR16ChoiceSl4
	SrsPeriodicityandoffsetR16ChoiceSl5
	SrsPeriodicityandoffsetR16ChoiceSl8
	SrsPeriodicityandoffsetR16ChoiceSl10
	SrsPeriodicityandoffsetR16ChoiceSl16
	SrsPeriodicityandoffsetR16ChoiceSl20
	SrsPeriodicityandoffsetR16ChoiceSl32
	SrsPeriodicityandoffsetR16ChoiceSl40
	SrsPeriodicityandoffsetR16ChoiceSl64
	SrsPeriodicityandoffsetR16ChoiceSl80
	SrsPeriodicityandoffsetR16ChoiceSl160
	SrsPeriodicityandoffsetR16ChoiceSl320
	SrsPeriodicityandoffsetR16ChoiceSl640
	SrsPeriodicityandoffsetR16ChoiceSl1280
	SrsPeriodicityandoffsetR16ChoiceSl2560
	SrsPeriodicityandoffsetR16ChoiceSl5120
	SrsPeriodicityandoffsetR16ChoiceSl10240
	SrsPeriodicityandoffsetR16ChoiceSl40960
	SrsPeriodicityandoffsetR16ChoiceSl81920
)

type SrsPeriodicityandoffsetR16 struct {
	Choice  uint64
	Sl1     *struct{}
	Sl2     *utils.INTEGER `lb:0,ub:1`
	Sl4     *utils.INTEGER `lb:0,ub:3`
	Sl5     *utils.INTEGER `lb:0,ub:4`
	Sl8     *utils.INTEGER `lb:0,ub:7`
	Sl10    *utils.INTEGER `lb:0,ub:9`
	Sl16    *utils.INTEGER `lb:0,ub:15`
	Sl20    *utils.INTEGER `lb:0,ub:19`
	Sl32    *utils.INTEGER `lb:0,ub:31`
	Sl40    *utils.INTEGER `lb:0,ub:39`
	Sl64    *utils.INTEGER `lb:0,ub:63`
	Sl80    *utils.INTEGER `lb:0,ub:79`
	Sl160   *utils.INTEGER `lb:0,ub:159`
	Sl320   *utils.INTEGER `lb:0,ub:319`
	Sl640   *utils.INTEGER `lb:0,ub:639`
	Sl1280  *utils.INTEGER `lb:0,ub:1279`
	Sl2560  *utils.INTEGER `lb:0,ub:2559`
	Sl5120  *utils.INTEGER `lb:0,ub:5119`
	Sl10240 *utils.INTEGER `lb:0,ub:10239`
	Sl40960 *utils.INTEGER `lb:0,ub:40959`
	Sl81920 *utils.INTEGER `lb:0,ub:81919`
}
