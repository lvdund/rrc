package ies

// FailureInfoRLC-Bearer ::= SEQUENCE
type FailureinforlcBearer struct {
	Cellgroupid            Cellgroupid
	Logicalchannelidentity Logicalchannelidentity
	Failuretype            FailureinforlcBearerFailuretype
}
