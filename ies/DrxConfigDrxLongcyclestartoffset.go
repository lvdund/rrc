package ies

import "rrc/utils"

// DRX-Config-drx-LongCycleStartOffset ::= CHOICE
const (
	DrxConfigDrxLongcyclestartoffsetChoiceNothing = iota
	DrxConfigDrxLongcyclestartoffsetChoiceMs10
	DrxConfigDrxLongcyclestartoffsetChoiceMs20
	DrxConfigDrxLongcyclestartoffsetChoiceMs32
	DrxConfigDrxLongcyclestartoffsetChoiceMs40
	DrxConfigDrxLongcyclestartoffsetChoiceMs60
	DrxConfigDrxLongcyclestartoffsetChoiceMs64
	DrxConfigDrxLongcyclestartoffsetChoiceMs70
	DrxConfigDrxLongcyclestartoffsetChoiceMs80
	DrxConfigDrxLongcyclestartoffsetChoiceMs128
	DrxConfigDrxLongcyclestartoffsetChoiceMs160
	DrxConfigDrxLongcyclestartoffsetChoiceMs256
	DrxConfigDrxLongcyclestartoffsetChoiceMs320
	DrxConfigDrxLongcyclestartoffsetChoiceMs512
	DrxConfigDrxLongcyclestartoffsetChoiceMs640
	DrxConfigDrxLongcyclestartoffsetChoiceMs1024
	DrxConfigDrxLongcyclestartoffsetChoiceMs1280
	DrxConfigDrxLongcyclestartoffsetChoiceMs2048
	DrxConfigDrxLongcyclestartoffsetChoiceMs2560
	DrxConfigDrxLongcyclestartoffsetChoiceMs5120
	DrxConfigDrxLongcyclestartoffsetChoiceMs10240
)

type DrxConfigDrxLongcyclestartoffset struct {
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
