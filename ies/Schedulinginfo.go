package ies

// SchedulingInfo ::= SEQUENCE
type Schedulinginfo struct {
	SiBroadcaststatus SchedulinginfoSiBroadcaststatus
	SiPeriodicity     SchedulinginfoSiPeriodicity
	SibMappinginfo    SibMapping
}
