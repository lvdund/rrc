package ies

// ReportConfigInterRAT-triggerType-event ::= SEQUENCE
// Extensible
type ReportconfiginterratTriggertypeEvent struct {
	Eventid       ReportconfiginterratTriggertypeEventEventid
	Hysteresis    Hysteresis
	Timetotrigger Timetotrigger
}
