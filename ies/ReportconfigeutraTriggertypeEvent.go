package ies

// ReportConfigEUTRA-triggerType-event ::= SEQUENCE
// Extensible
type ReportconfigeutraTriggertypeEvent struct {
	Eventid       ReportconfigeutraTriggertypeEventEventid
	Hysteresis    Hysteresis
	Timetotrigger Timetotrigger
}
