package ies

import "rrc/utils"

// SearchSpaceId ::= utils.INTEGER (0..maxNrofSearchSpaces-1)
type Searchspaceid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSearchSpaces1`
}
