package ies

import "rrc/utils"

// SIB9-timeInfo ::= SEQUENCE
type Sib9Timeinfo struct {
	Timeinfoutc        utils.INTEGER    `lb:0,ub:549755813887`
	Daylightsavingtime *utils.BITSTRING `lb:2,ub:2`
	Leapseconds        *utils.INTEGER   `lb:0,ub:128`
	Localtimeoffset    *utils.INTEGER   `lb:0,ub:64`
}
