package ies

import "rrc/utils"

// FreqsPriorityListGERAN ::= SEQUENCE OF FreqsPriorityGERAN
// SIZE (1..maxGNFG)
type Freqsprioritylistgeran struct {
	Value []Freqsprioritygeran
}
