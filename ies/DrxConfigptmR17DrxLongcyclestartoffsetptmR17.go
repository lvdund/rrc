package ies

import "rrc/utils"

// DRX-ConfigPTM-r17-drx-LongCycleStartOffsetPTM-r17 ::= CHOICE
const (
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceNothing = iota
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs10
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs20
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs32
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs40
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs60
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs64
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs70
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs80
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs128
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs160
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs256
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs320
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs512
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs640
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs1024
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs1280
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs2048
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs2560
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs5120
	DrxConfigptmR17DrxLongcyclestartoffsetptmR17ChoiceMs10240
)

type DrxConfigptmR17DrxLongcyclestartoffsetptmR17 struct {
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
