package ies

import "rrc/utils"

// SRS-PeriodicityAndOffset ::= CHOICE
const (
	SrsPeriodicityandoffsetChoiceNothing = iota
	SrsPeriodicityandoffsetChoiceSl1
	SrsPeriodicityandoffsetChoiceSl2
	SrsPeriodicityandoffsetChoiceSl4
	SrsPeriodicityandoffsetChoiceSl5
	SrsPeriodicityandoffsetChoiceSl8
	SrsPeriodicityandoffsetChoiceSl10
	SrsPeriodicityandoffsetChoiceSl16
	SrsPeriodicityandoffsetChoiceSl20
	SrsPeriodicityandoffsetChoiceSl32
	SrsPeriodicityandoffsetChoiceSl40
	SrsPeriodicityandoffsetChoiceSl64
	SrsPeriodicityandoffsetChoiceSl80
	SrsPeriodicityandoffsetChoiceSl160
	SrsPeriodicityandoffsetChoiceSl320
	SrsPeriodicityandoffsetChoiceSl640
	SrsPeriodicityandoffsetChoiceSl1280
	SrsPeriodicityandoffsetChoiceSl2560
)

type SrsPeriodicityandoffset struct {
	Choice uint64
	Sl1    *struct{}
	Sl2    *utils.INTEGER `lb:0,ub:1`
	Sl4    *utils.INTEGER `lb:0,ub:3`
	Sl5    *utils.INTEGER `lb:0,ub:4`
	Sl8    *utils.INTEGER `lb:0,ub:7`
	Sl10   *utils.INTEGER `lb:0,ub:9`
	Sl16   *utils.INTEGER `lb:0,ub:15`
	Sl20   *utils.INTEGER `lb:0,ub:19`
	Sl32   *utils.INTEGER `lb:0,ub:31`
	Sl40   *utils.INTEGER `lb:0,ub:39`
	Sl64   *utils.INTEGER `lb:0,ub:63`
	Sl80   *utils.INTEGER `lb:0,ub:79`
	Sl160  *utils.INTEGER `lb:0,ub:159`
	Sl320  *utils.INTEGER `lb:0,ub:319`
	Sl640  *utils.INTEGER `lb:0,ub:639`
	Sl1280 *utils.INTEGER `lb:0,ub:1279`
	Sl2560 *utils.INTEGER `lb:0,ub:2559`
}
