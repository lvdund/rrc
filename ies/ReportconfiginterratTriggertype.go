package ies

// ReportConfigInterRAT-triggerType ::= CHOICE
// Extensible
const (
	ReportconfiginterratTriggertypeChoiceNothing = iota
	ReportconfiginterratTriggertypeChoiceEvent
	ReportconfiginterratTriggertypeChoicePeriodical
)

type ReportconfiginterratTriggertype struct {
	Choice     uint64
	Event      *ReportconfiginterratTriggertypeEvent
	Periodical *ReportconfiginterratTriggertypePeriodical
}
