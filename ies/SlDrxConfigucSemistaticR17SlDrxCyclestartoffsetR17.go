package ies

import "rrc/utils"

// SL-DRX-ConfigUC-SemiStatic-r17-sl-drx-CycleStartOffset-r17 ::= CHOICE
const (
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceNothing = iota
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs10
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs20
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs32
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs40
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs60
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs64
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs70
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs80
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs128
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs160
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs256
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs320
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs512
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs640
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs1024
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs1280
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs2048
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs2560
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs5120
	SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17ChoiceMs10240
)

type SlDrxConfigucSemistaticR17SlDrxCyclestartoffsetR17 struct {
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
