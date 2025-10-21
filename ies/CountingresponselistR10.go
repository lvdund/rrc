package ies

import "rrc/utils"

// CountingResponseList-r10 ::= SEQUENCE OF CountingResponseInfo-r10
// SIZE (1..maxServiceCount)
type CountingresponselistR10 struct {
	Value []CountingresponseinfoR10
}
