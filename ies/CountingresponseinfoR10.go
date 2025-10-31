package ies

import "rrc/utils"

// CountingResponseInfo-r10 ::= SEQUENCE
// Extensible
type CountingresponseinfoR10 struct {
	CountingresponseserviceR10 utils.INTEGER `lb:0,ub:maxServiceCount1`
}
