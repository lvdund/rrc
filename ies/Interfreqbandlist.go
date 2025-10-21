package ies

import "rrc/utils"

// InterFreqBandList ::= SEQUENCE OF InterFreqBandInfo
// SIZE (1..maxBands)
type Interfreqbandlist struct {
	Value []Interfreqbandinfo
}
