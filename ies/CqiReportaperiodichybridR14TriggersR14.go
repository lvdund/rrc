package ies

// CQI-ReportAperiodicHybrid-r14-triggers-r14 ::= CHOICE
const (
	CqiReportaperiodichybridR14TriggersR14ChoiceNothing = iota
	CqiReportaperiodichybridR14TriggersR14ChoiceOnebitR14
	CqiReportaperiodichybridR14TriggersR14ChoiceTwobitR14
	CqiReportaperiodichybridR14TriggersR14ChoiceThreebitR14
)

type CqiReportaperiodichybridR14TriggersR14 struct {
	Choice      uint64
	OnebitR14   *CqiReportaperiodichybridR14TriggersR14OnebitR14
	TwobitR14   *CqiReportaperiodichybridR14TriggersR14TwobitR14
	ThreebitR14 *CqiReportaperiodichybridR14TriggersR14ThreebitR14
}
