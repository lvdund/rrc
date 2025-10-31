package ies

import "rrc/utils"

// SL-PSCCH-Config-r16 ::= SEQUENCE
// Extensible
type SlPscchConfigR16 struct {
	SlTimeresourcepscchR16 *SlPscchConfigR16SlTimeresourcepscchR16
	SlFreqresourcepscchR16 *SlPscchConfigR16SlFreqresourcepscchR16
	SlDmrsScrambleidR16    *utils.INTEGER `lb:0,ub:65535`
	SlNumreservedbitsR16   *utils.INTEGER `lb:0,ub:4`
}
