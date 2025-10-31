package ies

import "rrc/utils"

// FreqBandIndicatorNR ::= utils.INTEGER (1..1024)
type Freqbandindicatornr struct {
	Value utils.INTEGER `lb:0,ub:1024`
}
