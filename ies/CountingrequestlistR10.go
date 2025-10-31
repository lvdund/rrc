package ies

// CountingRequestList-r10 ::= SEQUENCE OF CountingRequestInfo-r10
// SIZE (1..maxServiceCount)
type CountingrequestlistR10 struct {
	Value []CountingrequestinfoR10 `lb:1,ub:maxServiceCount`
}
