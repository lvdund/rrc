package ies

import "rrc/utils"

// FreqBandIndicator ::= utils.INTEGER (1..maxFBI)
type Freqbandindicator struct {
	Value utils.INTEGER `lb:0,ub:maxFBI`
}
