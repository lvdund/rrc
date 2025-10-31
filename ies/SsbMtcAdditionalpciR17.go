package ies

import "rrc/utils"

// SSB-MTC-AdditionalPCI-r17 ::= SEQUENCE
type SsbMtcAdditionalpciR17 struct {
	AdditionalpciindexR17  AdditionalpciindexR17
	AdditionalpciR17       Physcellid
	PeriodicityR17         SsbMtcAdditionalpciR17PeriodicityR17
	SsbPositionsinburstR17 SsbMtcAdditionalpciR17SsbPositionsinburstR17
	SsPbchBlockpowerR17    utils.INTEGER `lb:0,ub:50`
}
