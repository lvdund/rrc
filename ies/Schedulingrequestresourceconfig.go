package ies

// SchedulingRequestResourceConfig ::= SEQUENCE
type Schedulingrequestresourceconfig struct {
	Schedulingrequestresourceid Schedulingrequestresourceid
	Schedulingrequestid         Schedulingrequestid
	Periodicityandoffset        *SchedulingrequestresourceconfigPeriodicityandoffset
	Resource                    *PucchResourceid
}
