package ies

import "rrc/utils"

// SystemInformationBlockType16-r11-timeInfo-r11 ::= SEQUENCE
type Systeminformationblocktype16R11TimeinfoR11 struct {
	TimeinfoutcR11        utils.INTEGER    `lb:0,ub:549755813887`
	DaylightsavingtimeR11 *utils.BITSTRING `lb:2,ub:2`
	LeapsecondsR11        *utils.INTEGER   `lb:0,ub:128`
	LocaltimeoffsetR11    *utils.INTEGER   `lb:0,ub:64`
}
