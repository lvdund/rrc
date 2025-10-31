package ies

// CQI-ReportConfig-r15 ::= CHOICE
const (
	CqiReportconfigR15ChoiceNothing = iota
	CqiReportconfigR15ChoiceRelease
	CqiReportconfigR15ChoiceSetup
)

type CqiReportconfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportconfigR15Setup
}
