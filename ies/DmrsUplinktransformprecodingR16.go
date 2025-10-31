package ies

import "rrc/utils"

// DMRS-UplinkTransformPrecoding-r16 ::= SEQUENCE
type DmrsUplinktransformprecodingR16 struct {
	Pi2bpskScramblingid0 *utils.INTEGER `lb:0,ub:65535`
	Pi2bpskScramblingid1 *utils.INTEGER `lb:0,ub:65535`
}
