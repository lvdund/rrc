package ies

// ReportConfigEUTRA-triggerType ::= CHOICE
// Extensible
const (
	ReportconfigeutraTriggertypeChoiceNothing = iota
	ReportconfigeutraTriggertypeChoiceEvent
	ReportconfigeutraTriggertypeChoicePeriodical
)

type ReportconfigeutraTriggertype struct {
	Choice     uint64
	Event      *ReportconfigeutraTriggertypeEvent
	Periodical *ReportconfigeutraTriggertypePeriodical
}
