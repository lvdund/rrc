package ies

import "rrc/utils"

// CountingRequestList-r10 ::= SEQUENCE OF CountingRequestInfo-r10
// SIZE (1..maxServiceCount)
type CountingrequestlistR10 struct {
	Value []CountingrequestinfoR10
}
