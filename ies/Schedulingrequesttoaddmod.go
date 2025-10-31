package ies

// SchedulingRequestToAddMod ::= SEQUENCE
type Schedulingrequesttoaddmod struct {
	Schedulingrequestid Schedulingrequestid
	SrProhibittimer     *SchedulingrequesttoaddmodSrProhibittimer
	SrTransmax          SchedulingrequesttoaddmodSrTransmax
}
