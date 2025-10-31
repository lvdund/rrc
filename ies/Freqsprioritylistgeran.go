package ies

// FreqsPriorityListGERAN ::= SEQUENCE OF FreqsPriorityGERAN
// SIZE (1..maxGNFG)
type Freqsprioritylistgeran struct {
	Value []Freqsprioritygeran `lb:1,ub:maxGNFG`
}
