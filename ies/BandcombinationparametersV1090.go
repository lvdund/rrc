package ies

import "rrc/utils"

// BandCombinationParameters-v1090 ::= SEQUENCE OF BandParameters-v1090
// SIZE (1..maxSimultaneousBands-r10)
type BandcombinationparametersV1090 struct {
	Value []BandparametersV1090
}
