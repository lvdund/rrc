package ies

import "rrc/utils"

// DRX-Info-drx-LongCycleStartOffset ::= CHOICE
const (
	DrxInfoDrxLongcyclestartoffsetChoiceNothing = iota
	DrxInfoDrxLongcyclestartoffsetChoiceMs10
	DrxInfoDrxLongcyclestartoffsetChoiceMs20
	DrxInfoDrxLongcyclestartoffsetChoiceMs32
	DrxInfoDrxLongcyclestartoffsetChoiceMs40
	DrxInfoDrxLongcyclestartoffsetChoiceMs60
	DrxInfoDrxLongcyclestartoffsetChoiceMs64
	DrxInfoDrxLongcyclestartoffsetChoiceMs70
	DrxInfoDrxLongcyclestartoffsetChoiceMs80
	DrxInfoDrxLongcyclestartoffsetChoiceMs128
	DrxInfoDrxLongcyclestartoffsetChoiceMs160
	DrxInfoDrxLongcyclestartoffsetChoiceMs256
	DrxInfoDrxLongcyclestartoffsetChoiceMs320
	DrxInfoDrxLongcyclestartoffsetChoiceMs512
	DrxInfoDrxLongcyclestartoffsetChoiceMs640
	DrxInfoDrxLongcyclestartoffsetChoiceMs1024
	DrxInfoDrxLongcyclestartoffsetChoiceMs1280
	DrxInfoDrxLongcyclestartoffsetChoiceMs2048
	DrxInfoDrxLongcyclestartoffsetChoiceMs2560
	DrxInfoDrxLongcyclestartoffsetChoiceMs5120
	DrxInfoDrxLongcyclestartoffsetChoiceMs10240
)

type DrxInfoDrxLongcyclestartoffset struct {
	Choice  uint64
	Ms10    *utils.INTEGER `lb:0,ub:9`
	Ms20    *utils.INTEGER `lb:0,ub:19`
	Ms32    *utils.INTEGER `lb:0,ub:31`
	Ms40    *utils.INTEGER `lb:0,ub:39`
	Ms60    *utils.INTEGER `lb:0,ub:59`
	Ms64    *utils.INTEGER `lb:0,ub:63`
	Ms70    *utils.INTEGER `lb:0,ub:69`
	Ms80    *utils.INTEGER `lb:0,ub:79`
	Ms128   *utils.INTEGER `lb:0,ub:127`
	Ms160   *utils.INTEGER `lb:0,ub:159`
	Ms256   *utils.INTEGER `lb:0,ub:255`
	Ms320   *utils.INTEGER `lb:0,ub:319`
	Ms512   *utils.INTEGER `lb:0,ub:511`
	Ms640   *utils.INTEGER `lb:0,ub:639`
	Ms1024  *utils.INTEGER `lb:0,ub:1023`
	Ms1280  *utils.INTEGER `lb:0,ub:1279`
	Ms2048  *utils.INTEGER `lb:0,ub:2047`
	Ms2560  *utils.INTEGER `lb:0,ub:2559`
	Ms5120  *utils.INTEGER `lb:0,ub:5119`
	Ms10240 *utils.INTEGER `lb:0,ub:10239`
}
